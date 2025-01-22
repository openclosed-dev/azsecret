package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type commandHandler struct {
	keyVaultName        string
	identity            string
	defaultKeyVaultName string
	defaultIdentity     string
}

func (h *commandHandler) prepare(cmd *cobra.Command, args []string) {

	if len(h.keyVaultName) == 0 {
		h.keyVaultName = h.defaultKeyVaultName
	}

	if len(h.identity) == 0 {
		h.identity = h.defaultIdentity
	}
}

func (h *commandHandler) handle(cmd *cobra.Command, args []string) error {

	var secretName = args[0]
	var client = newKeyVaultClient(h.keyVaultName)

	var err = client.authorize(h.identity)
	if err != nil {
		return err
	}

	secret, err := client.getSecret(secretName)
	if err != nil {
		return err
	}

	fmt.Print(secret)

	return nil
}

func execute() error {

	var handler = commandHandler{
		defaultKeyVaultName: strings.TrimSpace(os.Getenv("AZ_KEY_VAULT")),
		defaultIdentity:     strings.TrimSpace(os.Getenv("AZ_MANAGED_IDENTITY")),
	}

	var command = &cobra.Command{
		Use:    "azsecret [secret name]",
		Short:  "Retrieves a secret value stored in Azure Key Vault",
		Args:   cobra.ExactArgs(1),
		PreRun: handler.prepare,
		RunE:   handler.handle,
	}

	command.PersistentFlags().StringVarP(&handler.keyVaultName,
		"key-vault", "k", "",
		`Name of the Azure Key Vault.
Defaults to the value of environment variable AZ_KEY_VAULT.`)

	command.PersistentFlags().StringVarP(&handler.identity,
		"identity", "i", "",
		`Client ID of the Azure Managed Identity.
Defaults to the value of environment variable AZ_MANAGED_IDENTITY.`)

	if len(handler.defaultKeyVaultName) == 0 {
		command.MarkPersistentFlagRequired("key-vault")
	}

	return command.Execute()
}
