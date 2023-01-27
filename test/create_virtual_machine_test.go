package test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/FreekingDean/proxmox-api-go/proxmox"
	"github.com/FreekingDean/proxmox-api-go/proxmox/nodes/qemu"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateVirtualMachine(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		err := r.ParseForm()
		defer r.Body.Close()
		require.NoError(t, err)

		assert.Equal(t, "PVEAuthCookie=fake-cookie", r.Header.Get("Authorization"))
		assert.Equal(t, "fake-csrf", r.Header.Get("CSRFPreventionToken"))
		assert.Equal(t, "/nodes/some-node/qemu", r.URL.Path)
		assert.Equalf(t,
			strings.Split("file=test-file,bps=1,detect_zeroes=1,media=iso", ","),
			strings.Split(r.Form.Get("ide0"), ","),
			"Bad formvalue %+v", r.Form,
		)

		fmt.Fprintf(w, `{"data":"task-id"}`)
	}))
	c := proxmox.NewClient(testServer.URL)
	c.SetCookie("fake-cookie")
	c.SetCsrf("fake-csrf")
	q := qemu.New(c)
	req := qemu.CreateRequest{
		Node: "some-node",
		Ides: &qemu.IdenArr{
			qemu.Iden{
				File:         "test-file",
				Bps:          proxmox.Int(1),
				Media:        proxmox.String("iso"),
				DetectZeroes: proxmox.SpecialBool(true),
			},
		},
	}
	resp, err := q.Create(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "task-id", resp)
}
