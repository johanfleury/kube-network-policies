package networkpolicy

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/klog/v2"
	"sigs.k8s.io/kube-network-policies/pkg/api"
	"sigs.k8s.io/kube-network-policies/pkg/network"
)

func PacketLoggerFromContext(ctx context.Context, p *network.Packet, srcPod, dstPod *api.PodInfo) klog.Logger {
	return PacketLogger(klog.FromContext(ctx), p, srcPod, dstPod)
}

func PacketLogger(logger klog.Logger, p *network.Packet, srcPod, dstPod *api.PodInfo) klog.Logger {
	srcPodStr, dstPodStr := "external", "external"
	if srcPod != nil {
		srcPodStr = srcPod.Namespace.Name + "/" + srcPod.Name
	}
	if dstPod != nil {
		dstPodStr = dstPod.Namespace.Name + "/" + dstPod.Name
	}

	return logger.WithValues(
		"srcPod", srcPodStr,
		"dstPod", dstPodStr,
		"packet", p,
	)
}

// MatchesSelector returns true if the selector matches the given labels.
func MatchesSelector(selector *metav1.LabelSelector, lbls map[string]string) bool {
	s, err := metav1.LabelSelectorAsSelector(selector)
	if err != nil {
		klog.Errorf("error parsing label selector: %v", err)
		return false
	}
	return s.Matches(labels.Set(lbls))
}
