// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build sweep
// +build sweep

package codestarconnections

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/sweep"
)

func init() {
	resource.AddTestSweepers("aws_codestarconnections_connection", &resource.Sweeper{
		Name: "aws_codestarconnections_connection",
		F:    sweepConnections,
	})

	resource.AddTestSweepers("aws_codestarconnections_host", &resource.Sweeper{
		Name: "aws_codestarconnections_host",
		F:    sweepHosts,
		Dependencies: []string{
			"aws_codestarconnections_connection",
		},
	})
}

func sweepConnections(region string) error {
	ctx := sweep.Context(region)
	client, err := sweep.SharedRegionalSweepClient(ctx, region)
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}
	conn := client.CodeStarConnectionsClient(ctx)
	input := &codestarconnections.ListConnectionsInput{}
	sweepResources := make([]sweep.Sweepable, 0)

	pages := codestarconnections.NewListConnectionsPaginator(conn, input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx)

		if sweep.SkipSweepError(err) {
			log.Printf("[WARN] Skipping CodeStar Connections Connection sweep for %s: %s", region, err)
			return nil
		}

		if err != nil {
			return fmt.Errorf("error listing CodeStar Connections Connections (%s): %w", region, err)
		}

		for _, v := range page.Connections {
			r := resourceConnection()
			d := r.Data(nil)
			d.SetId(aws.ToString(v.ConnectionArn))

			sweepResources = append(sweepResources, sweep.NewSweepResource(r, d, client))
		}
	}

	err = sweep.SweepOrchestrator(ctx, sweepResources)

	if err != nil {
		return fmt.Errorf("error sweeping CodeStar Connections Connections (%s): %w", region, err)
	}

	return nil
}

func sweepHosts(region string) error {
	ctx := sweep.Context(region)
	client, err := sweep.SharedRegionalSweepClient(ctx, region)
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}
	conn := client.CodeStarConnectionsClient(ctx)
	input := &codestarconnections.ListHostsInput{}
	sweepResources := make([]sweep.Sweepable, 0)

	pages := codestarconnections.NewListHostsPaginator(conn, input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx)

		if sweep.SkipSweepError(err) {
			log.Printf("[WARN] Skipping CodeStar Connections Host sweep for %s: %s", region, err)
			return nil
		}

		if err != nil {
			return fmt.Errorf("error listing CodeStar Connections Hosts (%s): %w", region, err)
		}

		for _, v := range page.Hosts {
			r := resourceHost()
			d := r.Data(nil)
			d.SetId(aws.ToString(v.HostArn))

			sweepResources = append(sweepResources, sweep.NewSweepResource(r, d, client))
		}
	}

	err = sweep.SweepOrchestrator(ctx, sweepResources)

	if err != nil {
		return fmt.Errorf("error sweeping CodeStar Connections Hosts (%s): %w", region, err)
	}

	return nil
}
