/*
@Create Time     : 2020/9/10 16:08
@Project Name    : f5-rest-client
@File Name       : pool_stats_test.go
@Author          : zhangfeng
@Email           : 980252055@qq.com
@Software        : GoLand
*/

/*

 */
package ltm

import (
	"fmt"
	"github.com/e-XpertSolutions/f5-rest-client/f5"
	"log"
	"testing"
)

func TestPoolStatsResource_All(t *testing.T) {
	fmt.Println("sss")
	f5Client, err := f5.NewBasicClient("https://192.168.1.11", "admin", "admin")
	if err != nil {
		log.Fatal(err)
	}
	f5Client.DisableCertCheck()

	ltmClient := New(f5Client)
	fmt.Println(PoolStatsEndpoint)
	poolName := "zf"
	memberId := "192.168.0.100:27030"
	a, err := ltmClient.PoolStats().GetMemberStats(poolName, memberId)
	if err != nil {
		fmt.Println(err)
	}

	itemUrl := "https://localhost/mgmt/tm/ltm/pool/~Common~" + poolName + "/members/~Common~" + memberId + "/~Common~" + memberId + "/stats"
	MemberNestedStats := a.Entries[itemUrl]

	fmt.Println("---------------------")
	fmt.Println("Kind:", a.Kind)

	fmt.Println("---------------------")
	fmt.Println("SelfLink:", a.SelfLink)

	fmt.Println("---------------------")
	fmt.Println("Entries:", a.Entries)

	fmt.Println("---------------------")
	fmt.Println("MemberNestedStats:", MemberNestedStats)

	fmt.Println("---------------------")
	fmt.Println(MemberNestedStats.MemberNestedStats.Kind)

	fmt.Println("---------------------")
	fmt.Println(MemberNestedStats.MemberNestedStats.SelfLink)

	fmt.Println("---------------------")
	fmt.Println(MemberNestedStats.MemberNestedStats.Entries)

	fmt.Println("---------------------")
	fmt.Println(MemberNestedStats.MemberNestedStats.Entries.Addr.Description)

	fmt.Println("---------------------")
	fmt.Println(MemberNestedStats.MemberNestedStats.Entries.ServersideCurConns.Value)
}
