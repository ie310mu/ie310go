package redis_test

import (
	. "github.com/ie310mu/ie310go/forks/github.com/onsi/ginkgo"
	. "github.com/ie310mu/ie310go/forks/github.com/onsi/gomega"

	"github.com/ie310mu/ie310go/forks/github.com/go-redis/redis"
)

var _ = Describe("UniversalClient", func() {
	var client redis.UniversalClient

	AfterEach(func() {
		if client != nil {
			Expect(client.Close()).To(Succeed())
		}
	})

	It("should connect to failover servers", func() {
		client = redis.NewUniversalClient(&redis.UniversalOptions{
			MasterName: sentinelName,
			Addrs:      []string{":" + sentinelPort},
		})
		Expect(client.Ping().Err()).NotTo(HaveOccurred())
	})

	It("should connect to simple servers", func() {
		client = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs: []string{redisAddr},
		})
		Expect(client.Ping().Err()).NotTo(HaveOccurred())
	})

	It("should connect to clusters", func() {
		client = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs: cluster.addrs(),
		})
		Expect(client.Ping().Err()).NotTo(HaveOccurred())
	})

})
