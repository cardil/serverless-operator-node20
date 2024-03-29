// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	configv1 "github.com/openshift/api/config/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInfrastructures implements InfrastructureInterface
type FakeInfrastructures struct {
	Fake *FakeConfigV1
}

var infrastructuresResource = schema.GroupVersionResource{Group: "config.openshift.io", Version: "v1", Resource: "infrastructures"}

var infrastructuresKind = schema.GroupVersionKind{Group: "config.openshift.io", Version: "v1", Kind: "Infrastructure"}

// Get takes name of the infrastructure, and returns the corresponding infrastructure object, and an error if there is any.
func (c *FakeInfrastructures) Get(ctx context.Context, name string, options v1.GetOptions) (result *configv1.Infrastructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(infrastructuresResource, name), &configv1.Infrastructure{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Infrastructure), err
}

// List takes label and field selectors, and returns the list of Infrastructures that match those selectors.
func (c *FakeInfrastructures) List(ctx context.Context, opts v1.ListOptions) (result *configv1.InfrastructureList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(infrastructuresResource, infrastructuresKind, opts), &configv1.InfrastructureList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &configv1.InfrastructureList{ListMeta: obj.(*configv1.InfrastructureList).ListMeta}
	for _, item := range obj.(*configv1.InfrastructureList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested infrastructures.
func (c *FakeInfrastructures) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(infrastructuresResource, opts))
}

// Create takes the representation of a infrastructure and creates it.  Returns the server's representation of the infrastructure, and an error, if there is any.
func (c *FakeInfrastructures) Create(ctx context.Context, infrastructure *configv1.Infrastructure, opts v1.CreateOptions) (result *configv1.Infrastructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(infrastructuresResource, infrastructure), &configv1.Infrastructure{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Infrastructure), err
}

// Update takes the representation of a infrastructure and updates it. Returns the server's representation of the infrastructure, and an error, if there is any.
func (c *FakeInfrastructures) Update(ctx context.Context, infrastructure *configv1.Infrastructure, opts v1.UpdateOptions) (result *configv1.Infrastructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(infrastructuresResource, infrastructure), &configv1.Infrastructure{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Infrastructure), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInfrastructures) UpdateStatus(ctx context.Context, infrastructure *configv1.Infrastructure, opts v1.UpdateOptions) (*configv1.Infrastructure, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(infrastructuresResource, "status", infrastructure), &configv1.Infrastructure{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Infrastructure), err
}

// Delete takes name of the infrastructure and deletes it. Returns an error if one occurs.
func (c *FakeInfrastructures) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(infrastructuresResource, name, opts), &configv1.Infrastructure{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInfrastructures) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(infrastructuresResource, listOpts)

	_, err := c.Fake.Invokes(action, &configv1.InfrastructureList{})
	return err
}

// Patch applies the patch and returns the patched infrastructure.
func (c *FakeInfrastructures) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *configv1.Infrastructure, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(infrastructuresResource, name, pt, data, subresources...), &configv1.Infrastructure{})
	if obj == nil {
		return nil, err
	}
	return obj.(*configv1.Infrastructure), err
}
