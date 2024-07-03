package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
)

const BaseTfDirectory = "./terraform"

type TerraformOutput struct {
	Value string `json:"value"`
}

type ApplyOutput struct {
	InstanceID       string `json:"instance_id"`
	InstancePublicIP string `json:"instance_public_ip"`
}

func runCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = BaseTfDirectory
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running command %s %v: %v - output: %s", name, args, err, string(output))
	}
	return string(output), nil
}

func TerraformInitActivity(ctx context.Context) (string, error) {
	output, err := runCommand("terraform", "init", "-input=false")
	if err != nil {
		return "", err
	}
	return output, nil
}

func TerraformApplyActivity(ctx context.Context) (string, error) {
	output, err := runCommand("terraform", "apply", "-input=false", "-auto-approve")
	if err != nil {
		return "", err
	}
	return output, nil
}

func TerraformOutputActivity(ctx context.Context) (ApplyOutput, error) {
	outputValues, err := runCommand("terraform", "output", "-json")
	if err != nil {
		return ApplyOutput{}, err
	}

	var tfOutput map[string]TerraformOutput
	if err := json.Unmarshal([]byte(outputValues), &tfOutput); err != nil {
		return ApplyOutput{}, fmt.Errorf("error unmarshaling terraform output: %v", err)
	}

	applyOutput := ApplyOutput{
		InstanceID:       tfOutput["instance_id"].Value,
		InstancePublicIP: tfOutput["instance_public_ip"].Value,
	}

	return applyOutput, nil
}

// TODO ( Placeholder ) : Add more code to configure the instance post creation -
func ConfigureInstanceActivity(ctx context.Context, instancePublicIP string) error {
	// Here you can use instancePublicIP to SSH into the instance or perform any other configuration
	fmt.Printf("Configuring instance with IP: %s\n", instancePublicIP)
	// Implement the actual logic to configure the instance here
	return nil
}
