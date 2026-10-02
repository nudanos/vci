// SPDX-License-Identifier: MPL-2.0

package services

import "testing"

// Without systemd (a chroot, a container, the image build) Disable removes the
// wants symlink by hand and must not touch the D-Bus connection, which is nil.
func TestDisableWithoutSystemdDoesNotUseDBus(t *testing.T) {
	mgr := &Manager{available: false}
	if err := mgr.Disable("nudanos-test-component-that-is-not-enabled"); err != nil {
		t.Fatalf("Disable = %v, want nil (nothing to disable)", err)
	}
}
