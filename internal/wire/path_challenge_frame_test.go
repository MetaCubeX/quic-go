package wire

import (
	"io"

	"github.com/metacubex/quic-go/internal/protocol"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PATH_CHALLENGE frame", func() {
	Context("when parsing", func() {
		It("accepts sample frame", func() {
			b := []byte{1, 2, 3, 4, 5, 6, 7, 8}
			f, l, err := parsePathChallengeFrame(b, protocol.Version1)
			Expect(err).ToNot(HaveOccurred())
			Expect(f.Data).To(Equal([8]byte{1, 2, 3, 4, 5, 6, 7, 8}))
			Expect(l).To(Equal(len(b)))
		})

		It("errors on EOFs", func() {
			data := []byte{1, 2, 3, 4, 5, 6, 7, 8}
			_, l, err := parsePathChallengeFrame(data, protocol.Version1)
			Expect(err).ToNot(HaveOccurred())
			Expect(l).To(Equal(len(data)))
			for i := range data {
				_, _, err := parsePathChallengeFrame(data[:i], protocol.Version1)
				Expect(err).To(MatchError(io.EOF))
			}
		})
	})

	Context("when writing", func() {
		It("writes a sample frame", func() {
			frame := PathChallengeFrame{Data: [8]byte{0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe, 0x13, 0x37}}
			b, err := frame.Append(nil, protocol.Version1)
			Expect(err).ToNot(HaveOccurred())
			Expect(b).To(Equal([]byte{pathChallengeFrameType, 0xde, 0xad, 0xbe, 0xef, 0xca, 0xfe, 0x13, 0x37}))
		})

		It("has the correct length", func() {
			frame := PathChallengeFrame{}
			Expect(frame.Length(protocol.Version1)).To(Equal(protocol.ByteCount(9)))
		})
	})
})
