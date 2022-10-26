package main_test

import (
	main "a21hc3NpZ25tZW50"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var CreateChannel = func() (chan main.RowData, chan error) {
	ch := make(chan main.RowData)
	chErr := make(chan error)
	return ch, chErr
}

const errTimeoutMessage = "Timeout: execute test take time more than 250 miliseconds"

var _ = Describe("ProcessGetTLD", func() {
	When("field Domain is empty string", func() {
		It("should send error to channel with message 'domain name is empty'", func() {
			ch, errCh := CreateChannel()

			go main.ProcessGetTLD(main.RowData{
				RankWebsite: 1,
				Domain:      "",
				Valid:       true,
				RefIPs:      100,
			}, ch, errCh)

			select {
			case <-time.After(250 * time.Millisecond):
				Fail(errTimeoutMessage)
			case err := <-errCh:
				Expect(err).ToNot(BeNil())
				Expect(err).To(MatchError("domain name is empty"))
				Expect(err.Error()).To(Equal("domain name is empty"))
			}
		})
	})

	When("field Valid is false", func() {
		It("should send error to channel with message 'domain not valid'", func() {
			ch, errCh := CreateChannel()

			go main.ProcessGetTLD(main.RowData{
				RankWebsite: 1,
				Domain:      "google.com",
				Valid:       false,
				RefIPs:      10,
			}, ch, errCh)

			select {
			case <-time.After(250 * time.Millisecond):
				Fail(errTimeoutMessage)
			case err := <-errCh:
				Expect(err).ToNot(BeNil())
				Expect(err).To(MatchError("domain not valid"))
				Expect(err.Error()).To(Equal("domain not valid"))
			}

		})
	})

	When("field RefIps is -1", func() {
		It("should send error to channel with message 'domain RefIPs not valid'", func() {
			ch, errCh := CreateChannel()

			go main.ProcessGetTLD(main.RowData{
				RankWebsite: 1,
				Domain:      "google.com",
				Valid:       true,
				RefIPs:      -1,
			}, ch, errCh)

			select {
			case <-time.After(250 * time.Millisecond):
				Fail(errTimeoutMessage)
			case err := <-errCh:
				Expect(err).ToNot(BeNil())
				Expect(err).To(MatchError("domain RefIPs not valid"))
				Expect(err.Error()).To(Equal("domain RefIPs not valid"))
			}
		})
	})

	When("field Domain is not empty string, field Valid is true, and field RefIps is not -1", func() {
		It("should send data to channel", func() {
			ch, errCh := CreateChannel()

			go main.ProcessGetTLD(main.RowData{
				RankWebsite: 1,
				Domain:      "google.com",
				Valid:       true,
				RefIPs:      10,
			}, ch, errCh)

			select {
			case <-time.After(250 * time.Millisecond):
				Fail(errTimeoutMessage)
			case data := <-ch:
				Expect(data).ToNot(BeNil())
				Expect(data).To(Equal(main.RowData{
					RankWebsite: 1,
					Domain:      "google.com",
					Valid:       true,
					RefIPs:      10,
					TLD:         ".com",
					IDN_TLD:     ".co.id",
				}))

			}
		})
	})
})

var testData = []main.RowData{
	{RankWebsite: 1, Domain: "google.com", Valid: true, RefIPs: 2404064},
	{RankWebsite: 2, Domain: "facebook.com", Valid: true, RefIPs: 2547862},
	{RankWebsite: 3, Domain: "youtube.com", Valid: true, RefIPs: 2067945},
	{RankWebsite: 4, Domain: "twitter.com", Valid: true, RefIPs: 2046616},
	{RankWebsite: 5, Domain: "instagram.com", Valid: true, RefIPs: 1694854},
	{RankWebsite: 6, Domain: "linkedin.com", Valid: true, RefIPs: 1426379},
	{RankWebsite: 7, Domain: "apple.com", Valid: true, RefIPs: 908289},
	{RankWebsite: 8, Domain: "microsoft.com", Valid: true, RefIPs: 823943},
	{RankWebsite: 9, Domain: "wikipedia.org", Valid: true, RefIPs: 962648},
	{RankWebsite: 10, Domain: "wordpress.org", Valid: true, RefIPs: 1073003},
	{RankWebsite: 11, Domain: "googletagmanager.com", Valid: true, RefIPs: 865833},
	{RankWebsite: 12, Domain: "youtu.be", Valid: true, RefIPs: 783542},
	{RankWebsite: 13, Domain: "en.wikipedia.org", Valid: true, RefIPs: 720304},
	{RankWebsite: 14, Domain: "pinterest.com", Valid: true, RefIPs: 851879},
	{RankWebsite: 15, Domain: "play.google.com", Valid: true, RefIPs: 606681},
	{RankWebsite: 16, Domain: "vimeo.com", Valid: true, RefIPs: 724184},
	{RankWebsite: 17, Domain: "github.com", Valid: true, RefIPs: 526163},
	{RankWebsite: 18, Domain: "maps.google.com", Valid: true, RefIPs: 697449},
	{RankWebsite: 19, Domain: "plus.google.com", Valid: true, RefIPs: 716411},
	{RankWebsite: 20, Domain: "goo.gl", Valid: true, RefIPs: 640887},
	{RankWebsite: 21, Domain: "bit.ly", Valid: true, RefIPs: 568027},
	{RankWebsite: 22, Domain: "adobe.com", Valid: true, RefIPs: 515112},
	{RankWebsite: 23, Domain: "docs.google.com", Valid: true, RefIPs: 533835},
	{RankWebsite: 24, Domain: "wordpress.com", Valid: true, RefIPs: 575098},
	{RankWebsite: 25, Domain: "blogspot.com", Valid: true, RefIPs: 539018},
	{RankWebsite: 26, Domain: "amazon.com", Valid: true, RefIPs: 510322},
	{RankWebsite: 27, Domain: "player.vimeo.com", Valid: true, RefIPs: 563305},
	{RankWebsite: 28, Domain: "itunes.apple.com", Valid: true, RefIPs: 429682},
	{RankWebsite: 29, Domain: "mozilla.org", Valid: true, RefIPs: 454323},
	{RankWebsite: 30, Domain: "macromedia.com", Valid: true, RefIPs: 468646},
	{RankWebsite: 31, Domain: "apps.apple.com", Valid: true, RefIPs: 366239},
	{RankWebsite: 32, Domain: "drive.google.com", Valid: true, RefIPs: 414206},
	{RankWebsite: 33, Domain: "yahoo.com", Valid: true, RefIPs: 383452},
	{RankWebsite: 34, Domain: "whatsapp.com", Valid: true, RefIPs: 494252},
	{RankWebsite: 35, Domain: "lenyar.ru", Valid: true, RefIPs: 413315},
	{RankWebsite: 36, Domain: "download.macromedia.com", Valid: true, RefIPs: 400577},
	{RankWebsite: 37, Domain: "tumblr.com", Valid: true, RefIPs: 442738},
	{RankWebsite: 38, Domain: "europa.eu", Valid: true, RefIPs: 437416},
	{RankWebsite: 39, Domain: "flickr.com", Valid: true, RefIPs: 354123},
	{RankWebsite: 40, Domain: "api.whatsapp.com", Valid: true, RefIPs: 432601},
	{RankWebsite: 41, Domain: "reddit.com", Valid: true, RefIPs: 374387},
	{RankWebsite: 42, Domain: "nih.gov", Valid: true, RefIPs: 321773},
	{RankWebsite: 43, Domain: "qq.com", Valid: true, RefIPs: 741568},
	{RankWebsite: 44, Domain: "gravatar.com", Valid: true, RefIPs: 420611},
	{RankWebsite: 45, Domain: "support.microsoft.com", Valid: true, RefIPs: 303121},
	{RankWebsite: 46, Domain: "nytimes.com", Valid: true, RefIPs: 322893},
	{RankWebsite: 47, Domain: "amazonaws.com", Valid: true, RefIPs: 269143},
	{RankWebsite: 48, Domain: "apache.org", Valid: true, RefIPs: 228407},
	{RankWebsite: 49, Domain: "policies.google.com", Valid: true, RefIPs: 376694},
	{RankWebsite: 50, Domain: "sites.google.com", Valid: true, RefIPs: 283257},
	{RankWebsite: 51, Domain: "github.io", Valid: true, RefIPs: 252770},
	{RankWebsite: 52, Domain: "soundcloud.com", Valid: true, RefIPs: 324764},
	{RankWebsite: 53, Domain: "w3.org", Valid: true, RefIPs: 234732},
	{RankWebsite: 54, Domain: "t.co", Valid: true, RefIPs: 311962},
	{RankWebsite: 55, Domain: "medium.com", Valid: true, RefIPs: 273490},
	{RankWebsite: 56, Domain: "forbes.com", Valid: true, RefIPs: 277123},
	{RankWebsite: 57, Domain: "ec.europa.eu", Valid: true, RefIPs: 332359},
	{RankWebsite: 58, Domain: "forms.gle", Valid: true, RefIPs: 260086},
	{RankWebsite: 59, Domain: "miit.gov.cn", Valid: true, RefIPs: 810196},
	{RankWebsite: 60, Domain: "zoom.us", Valid: true, RefIPs: 251265},
	{RankWebsite: 61, Domain: "spotify.com", Valid: true, RefIPs: 291982},
	{RankWebsite: 62, Domain: "baidu.com", Valid: true, RefIPs: 550260},
	{RankWebsite: 63, Domain: "www.ncbi.nlm.nih.gov", Valid: true, RefIPs: 260034},
	{RankWebsite: 64, Domain: "archive.org", Valid: true, RefIPs: 259081},
	{RankWebsite: 65, Domain: "theguardian.com", Valid: true, RefIPs: 266709},
	{RankWebsite: 66, Domain: "creativecommons.org", Valid: true, RefIPs: 236048},
	{RankWebsite: 67, Domain: "sourceforge.net", Valid: true, RefIPs: 205478},
	{RankWebsite: 68, Domain: "beian.miit.gov.cn", Valid: true, RefIPs: 720415},
	{RankWebsite: 69, Domain: "t.me", Valid: true, RefIPs: 300511},
	{RankWebsite: 70, Domain: "cnn.com", Valid: true, RefIPs: 250416},
	{RankWebsite: 71, Domain: "dropbox.com", Valid: true, RefIPs: 232091},
	{RankWebsite: 72, Domain: "who.int", Valid: true, RefIPs: 219883},
	{RankWebsite: 73, Domain: "cloudflare.com", Valid: true, RefIPs: 274188},
	{RankWebsite: 74, Domain: "wixsite.com", Valid: true, RefIPs: 237191},
	{RankWebsite: 75, Domain: "open.spotify.com", Valid: true, RefIPs: 260963},
	{RankWebsite: 76, Domain: "wa.me", Valid: true, RefIPs: 301259},
	{RankWebsite: 77, Domain: "bbc.co.uk", Valid: true, RefIPs: 236902},
	{RankWebsite: 78, Domain: "paypal.com", Valid: true, RefIPs: 259460},
	{RankWebsite: 79, Domain: "m.facebook.com", Valid: true, RefIPs: 251505},
	{RankWebsite: 80, Domain: "office.com", Valid: true, RefIPs: 202871},
	{RankWebsite: 81, Domain: "issuu.com", Valid: true, RefIPs: 215445},
	{RankWebsite: 82, Domain: "weebly.com", Valid: true, RefIPs: 217433},
	{RankWebsite: 83, Domain: "cdc.gov", Valid: true, RefIPs: 222642},
	{RankWebsite: 84, Domain: "bbc.com", Valid: true, RefIPs: 217200},
	{RankWebsite: 85, Domain: "sciencedirect.com", Valid: true, RefIPs: 208425},
	{RankWebsite: 86, Domain: "httpd.apache.org", Valid: true, RefIPs: 160261},
	{RankWebsite: 87, Domain: "vk.com", Valid: true, RefIPs: 284952},
	{RankWebsite: 88, Domain: "tinyurl.com", Valid: true, RefIPs: 203408},
	{RankWebsite: 89, Domain: "reuters.com", Valid: true, RefIPs: 202723},
	{RankWebsite: 90, Domain: "accounts.google.com", Valid: true, RefIPs: 187353},
	{RankWebsite: 91, Domain: "tiktok.com", Valid: true, RefIPs: 221003},
	{RankWebsite: 92, Domain: "nginx.org", Valid: true, RefIPs: 172682},
	{RankWebsite: 93, Domain: "opera.com", Valid: true, RefIPs: 214342},
	{RankWebsite: 94, Domain: "washingtonpost.com", Valid: true, RefIPs: 209202},
	{RankWebsite: 95, Domain: "web.archive.org", Valid: true, RefIPs: 201174},
	{RankWebsite: 96, Domain: "youtube-nocookie.com", Valid: true, RefIPs: 207717},
	{RankWebsite: 97, Domain: "oracle.com", Valid: true, RefIPs: 154311},
	{RankWebsite: 98, Domain: "bloomberg.com", Valid: true, RefIPs: 192359},
	{RankWebsite: 99, Domain: "live.com", Valid: true, RefIPs: 188214},
	{RankWebsite: 100, Domain: "imdb.com", Valid: true, RefIPs: 211304},
}

var _ = Describe("FilterAndGetDomain", func() {
	ReportAfterEach(func(report SpecReport) {
		if report.RunTime > (time.Millisecond * 300) {
			AbortSuite(errTimeoutMessage)
		}
	})

	Context("error cases", func() {
		When("one of list data have empty domain", func() {
			It("should return error with message 'domain name is empty'", func() {
				var test = make([]main.RowData, len(testData))
				copy(test, testData)

				test[30].Domain = ""

				_, err := main.FilterAndFillData(".com", test)

				Expect(err).ToNot(BeNil())
				Expect(err).To(MatchError("domain name is empty"))
				Expect(err.Error()).To(Equal("domain name is empty"))
			})
		})

		When("one of list data have valid is false", func() {
			It("should return error with message 'domain not valid'", func() {
				var test = make([]main.RowData, len(testData))
				copy(test, testData)
				test[20].Valid = false

				_, err := main.FilterAndFillData(".com", test)

				Expect(err).ToNot(BeNil())
				Expect(err).To(MatchError("domain not valid"))
				Expect(err.Error()).To(Equal("domain not valid"))
			})
		})

		When("one of list data have RefIPs -1", func() {
			It("should return error with message 'domain RefIPs not valid'", func() {
				var test = make([]main.RowData, len(testData))
				copy(test, testData)
				test[10].RefIPs = -1

				_, err := main.FilterAndFillData(".com", test)

				Expect(err).ToNot(BeNil())
				Expect(err).To(MatchError("domain RefIPs not valid"))
				Expect(err.Error()).To(Equal("domain RefIPs not valid"))
			})
		})
	})

	Context("Concurrent process", func() {
		When("total data execute is 50", func() {
			It("should execute less than 200 millisecond", func() {
				main.FuncProcessGetTLD = func(data main.RowData, ch chan main.RowData, errCh chan error) {
					time.Sleep(100 * time.Millisecond)
					main.ProcessGetTLD(data, ch, errCh)
				}

				start := time.Now()
				_, err := main.FilterAndFillData(".com", testData[:50])

				Expect(err).To(BeNil())
				Expect(time.Since(start)).To(BeNumerically("<", 200*time.Millisecond))
			})
		})

		When("total data execute is 75", func() {
			It("should execute less than 200 millisecond", func() {
				main.FuncProcessGetTLD = func(data main.RowData, ch chan main.RowData, errCh chan error) {
					time.Sleep(100 * time.Millisecond)
					main.ProcessGetTLD(data, ch, errCh)
				}

				start := time.Now()
				_, err := main.FilterAndFillData(".com", testData[:75])

				Expect(err).To(BeNil())
				Expect(time.Since(start)).To(BeNumerically("<", 200*time.Millisecond))
			})
		})
	})

	Context("success cases", func() {
		When("TLD input is .com", func() {
			It("should return filter list data with TLD is .com", func() {
				result, err := main.FilterAndFillData(".com", testData)

				Expect(err).To(BeNil())
				Expect(len(result)).To(Equal(69))
			})
		})

		When("TLD input is .co.id", func() {
			It("should return filter list data with TLD is .co.id", func() {
				result, err := main.FilterAndFillData(".co.id", testData)

				Expect(err).To(BeNil())
				Expect(len(result)).To(Equal(0))
			})
		})

		When("TLD input is .org", func() {
			It("should return filter list data with TLD is .org", func() {
				result, err := main.FilterAndFillData(".org", testData)

				Expect(err).To(BeNil())
				Expect(len(result)).To(Equal(11))
			})
		})
	})
})
