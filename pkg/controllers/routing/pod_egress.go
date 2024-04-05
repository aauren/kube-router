package routing

import (
	"errors"
	"fmt"
	"time"

	"github.com/cloudnativelabs/kube-router/pkg/metrics"

	"k8s.io/klog/v2"
)

// set up MASQUERADE rule so that egress traffic from the pods gets masqueraded to node's IP

var (
	podEgressArgs4 = []string{"-m", "set", "--match-set", podSubnetsIPSetName, "src",
		"-m", "set", "!", "--match-set", podSubnetsIPSetName, "dst",
		"-m", "set", "!", "--match-set", nodeAddrsIPSetName, "dst",
		"-j", "MASQUERADE"}
	podEgressArgs6 = []string{"-m", "set", "--match-set", "inet6:" + podSubnetsIPSetName, "src",
		"-m", "set", "!", "--match-set", "inet6:" + podSubnetsIPSetName, "dst",
		"-m", "set", "!", "--match-set", "inet6:" + nodeAddrsIPSetName, "dst",
		"-j", "MASQUERADE"}
	podEgressArgsBad4 = [][]string{{"-m", "set", "--match-set", podSubnetsIPSetName, "src",
		"-m", "set", "!", "--match-set", podSubnetsIPSetName, "dst",
		"-j", "MASQUERADE"}}
	podEgressArgsBad6 = [][]string{{"-m", "set", "--match-set", "inet6:" + podSubnetsIPSetName, "src",
		"-m", "set", "!", "--match-set", "inet6:" + podSubnetsIPSetName, "dst",
		"-j", "MASQUERADE"}}
)

func (nrc *NetworkRoutingController) createPodEgressRule() error {
	iptablesCmdHandler, err := nrc.newIptablesCmdHandler()
	if err != nil {
		return errors.New("Failed create iptables handler:" + err.Error())
	}

	podEgressArgs := podEgressArgs4
	if nrc.isIpv6 {
		podEgressArgs = podEgressArgs6
	}

	cmdStart := time.Now()
	defer func() {
		cmdTime := time.Since(cmdStart)
		if nrc.MetricsEnabled {
			metrics.ControllerRoutesIPTablesCommandTime.Observe(cmdTime.Seconds())
		}
		klog.V(2).Infof("IPTables command time took: %v", cmdTime)
	}()

	if iptablesCmdHandler.HasRandomFully() {
		podEgressArgs = append(podEgressArgs, "--random-fully")
	}

	err = iptablesCmdHandler.AppendUnique("nat", "POSTROUTING", podEgressArgs...)
	if err != nil {
		return errors.New("Failed to add iptables rule to masquerade outbound traffic from pods: " +
			err.Error() + "External connectivity will not work.")

	}

	klog.V(1).Infof("Added iptables rule to masquerade outbound traffic from pods.")
	return nil
}

func (nrc *NetworkRoutingController) deletePodEgressRule() error {
	iptablesCmdHandler, err := nrc.newIptablesCmdHandler()
	if err != nil {
		return errors.New("Failed create iptables handler:" + err.Error())
	}

	podEgressArgs := podEgressArgs4
	if nrc.isIpv6 {
		podEgressArgs = podEgressArgs6
	}

	cmdStart := time.Now()
	defer func() {
		cmdTime := time.Since(cmdStart)
		if nrc.MetricsEnabled {
			metrics.ControllerRoutesIPTablesCommandTime.Observe(cmdTime.Seconds())
		}
		klog.V(2).Infof("IPTables command time took: %v", cmdTime)
	}()

	if iptablesCmdHandler.HasRandomFully() {
		podEgressArgs = append(podEgressArgs, "--random-fully")
	}

	exists, err := iptablesCmdHandler.Exists("nat", "POSTROUTING", podEgressArgs...)
	if err != nil {
		return errors.New("Failed to lookup iptables rule to masquerade outbound traffic from pods: " + err.Error())
	}

	if exists {
		err = iptablesCmdHandler.Delete("nat", "POSTROUTING", podEgressArgs...)
		if err != nil {
			return errors.New("Failed to delete iptables rule to masquerade outbound traffic from pods: " +
				err.Error() + ". Pod egress might still work...")
		}
		klog.Infof("Deleted iptables rule to masquerade outbound traffic from pods.")
	}

	return nil
}

func (nrc *NetworkRoutingController) deleteBadPodEgressRules() error {
	iptablesCmdHandler, err := nrc.newIptablesCmdHandler()
	if err != nil {
		return errors.New("Failed create iptables handler:" + err.Error())
	}
	podEgressArgsBad := podEgressArgsBad4
	if nrc.isIpv6 {
		podEgressArgsBad = podEgressArgsBad6
	}

	cmdStart := time.Now()
	defer func() {
		cmdTime := time.Since(cmdStart)
		if nrc.MetricsEnabled {
			metrics.ControllerRoutesIPTablesCommandTime.Observe(cmdTime.Seconds())
		}
		klog.V(2).Infof("IPTables command time took: %v", cmdTime)
	}()

	// If random fully is supported remove the original rule as well
	if iptablesCmdHandler.HasRandomFully() {
		if !nrc.isIpv6 {
			podEgressArgsBad = append(podEgressArgsBad, podEgressArgs4)
		} else {
			podEgressArgsBad = append(podEgressArgsBad, podEgressArgs6)
		}
	}

	for _, args := range podEgressArgsBad {
		exists, err := iptablesCmdHandler.Exists("nat", "POSTROUTING", args...)
		if err != nil {
			return fmt.Errorf("failed to lookup iptables rule: %s", err.Error())
		}

		if exists {
			err = iptablesCmdHandler.Delete("nat", "POSTROUTING", args...)
			if err != nil {
				return fmt.Errorf("failed to delete old/bad iptables rule to masquerade outbound traffic "+
					"from pods: %s. Pod egress might still work, or bugs may persist after upgrade", err)
			}
			klog.Infof("Deleted old/bad iptables rule to masquerade outbound traffic from pods.")
		}
	}

	return nil
}
