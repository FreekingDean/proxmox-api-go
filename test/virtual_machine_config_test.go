package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FreekingDean/proxmox-api-go/proxmox"
	"github.com/FreekingDean/proxmox-api-go/proxmox/nodes/qemu"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVmConfig(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		err := r.ParseForm()
		defer r.Body.Close()
		require.NoError(t, err)

		assert.Equal(t, "PVEAuthCookie=fake-cookie", r.Header.Get("Authorization"))
		assert.Equal(t, "/nodes/some-node/qemu/101/config", r.URL.Path)
		fmt.Fprintf(w, `{"data":{"scsi0": "testfile,aio=native,ssd=1"}}`)
	}))
	c := proxmox.NewClient(testServer.URL)
	c.SetCookie("fake-cookie")
	q := qemu.New(c)
	req := qemu.VmConfigRequest{
		Node: "some-node",
		Vmid: 101,
	}
	resp, err := q.VmConfig(context.Background(), req)
	if assert.NoError(t, err) {
		if assert.NotNil(t, resp.Scsis) && assert.Equal(t, 1, len(*resp.Scsis)) {
			assert.Equal(t, "testfile", (*resp.Scsis)[0].File)
			if assert.NotNil(t, (*resp.Scsis)[0].Ssd) {
				assert.Equal(t, true, bool(*(*resp.Scsis)[0].Ssd))
			}
			if assert.NotNil(t, (*resp.Scsis)[0].Aio) {
				assert.Equal(t, qemu.ScsiAio_NATIVE, *(*resp.Scsis)[0].Aio)
			}
		}
	}
}
