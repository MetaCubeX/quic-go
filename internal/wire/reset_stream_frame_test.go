package wire

import (
	"github.com/metacubex/quic-go/internal/protocol"
	"github.com/metacubex/quic-go/internal/qerr"
	"github.com/metacubex/quic-go/quicvarint"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RESET_STREAM frame", func() {
	Context("when parsing", func() {
		It("accepts sample frame", func() {
			data := encodeVarInt(0xdeadbeef)                  // stream ID
			data = append(data, encodeVarInt(0x1337)...)      // error code
			data = append(data, encodeVarInt(0x987654321)...) // byte offset
			frame, l, err := parseResetStreamFrame(data, protocol.Version1)
			Expect(err).ToNot(HaveOccurred())
			Expect(frame.StreamID).To(Equal(protocol.StreamID(0xdeadbeef)))
			Expect(frame.FinalSize).To(Equal(protocol.ByteCount(0x987654321)))
			Expect(frame.ErrorCode).To(Equal(qerr.StreamErrorCode(0x1337)))
			Expect(l).To(Equal(len(data)))
		})

		It("errors on EOFs", func() {
			data := encodeVarInt(0xdeadbeef)                  // stream ID
			data = append(data, encodeVarInt(0x1337)...)      // error code
			data = append(data, encodeVarInt(0x987654321)...) // byte offset
			_, l, err := parseResetStreamFrame(data, protocol.Version1)
			Expect(err).NotTo(HaveOccurred())
			Expect(l).To(Equal(len(data)))
			for i := range data {
				_, _, err := parseResetStreamFrame(data[:i], protocol.Version1)
				Expect(err).To(HaveOccurred())
			}
		})
	})

	Context("when writing", func() {
		It("writes a sample frame", func() {
			frame := ResetStreamFrame{
				StreamID:  0x1337,
				FinalSize: 0x11223344decafbad,
				ErrorCode: 0xcafe,
			}
			b, err := frame.Append(nil, protocol.Version1)
			Expect(err).ToNot(HaveOccurred())
			expected := []byte{resetStreamFrameType}
			expected = append(expected, encodeVarInt(0x1337)...)
			expected = append(expected, encodeVarInt(0xcafe)...)
			expected = append(expected, encodeVarInt(0x11223344decafbad)...)
			Expect(b).To(Equal(expected))
		})

		It("has the correct length", func() {
			rst := ResetStreamFrame{
				StreamID:  0x1337,
				FinalSize: 0x1234567,
				ErrorCode: 0xde,
			}
			expectedLen := 1 + quicvarint.Len(0x1337) + quicvarint.Len(0x1234567) + 2
			Expect(rst.Length(protocol.Version1)).To(BeEquivalentTo(expectedLen))
		})
	})
})
