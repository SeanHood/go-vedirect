package vedirect

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStreamHEXFrame(t *testing.T) {
	//assert := assert.New(t)

	// Convert []bytes to an io.Reader type
	testFrameReader := bytes.NewReader([]byte(":foo\n"))

	s := NewStream(testFrameReader)

	//var block Block
	// var checksum int

	// for checksum == 0 {
	// 	_, checksum = s.ReadBlock()
	// 	break
	// }

	//s.ReadBlock()
	assert.Equal(t, 0, s.State)

	//s.ReadBlock()

	//assert.Equal(1, s.State)

}

func TestNewStreamValidFrame(t *testing.T) {
	assert := assert.New(t)

	// Convert []bytes to an io.Reader type
	testFrameReader := bytes.NewReader(testNewStreamValidChecksum())

	s := NewStream(testFrameReader)

	var block Block
	var checksum int

	for checksum == 0 {
		block, checksum = s.ReadBlock()
		break
	}

	fields := block.Fields()

	// Test our getter returns
	assert.Equal(block.fields, block.Fields())

	// Test a couple feild's we're expecting
	assert.Equal("146", fields["FW"])
	assert.Equal("31300", fields["VPV"])

	// Expect t the same nummber of feilds in as we get out
	assert.Equal(20, len(fields))

}

func TestNewStreamInvalidChecksum(t *testing.T) {
	assert := assert.New(t)

	// Convert []bytes to an io.Reader type
	testFrameReader := bytes.NewReader(testNewStreamInvalidChecksum())

	s := NewStream(testFrameReader)

	var frame Block
	var checksum int

	for checksum == 0 {
		_, checksum = s.ReadBlock()
		break
	}

	// Checksum should not be 0, as it's invalid
	assert.NotEqual(0, checksum)

	// We're expecting an empty frame since the checksum wasn't valid
	assert.Equal(map[string]string(map[string]string(nil)), frame.fields)

}

func testNewStreamValidChecksum() []byte {

	return []byte("\r\n" +
		"PID\t0xA053\r\n" +
		"FW\t146\r\n" +
		"SER#\t12288FHJQXX\r\n" +
		"V\t13210\r\n" +
		"I\t2360\r\n" +
		"VPV\t31300\r\n" +
		"PPV\t37\r\n" +
		"CS\t3\r\n" +
		"MPPT\t2\r\n" +
		"OR\t0x00000000\r\n" +
		"ERR\t0\r\n" +
		"LOAD\tON\r\n" +
		"IL\t400\r\n" +
		"H19\t290\r\n" +
		"H20\t20\r\n" +
		"H21\t46\r\n" +
		"H22\t0\r\n" +
		"H23\t0\r\n" +
		"HSDS\t42\r\n" +
		"Checksum\tT\r\n")

}

func testNewStreamInvalidChecksum() []byte {

	return []byte("\r\n" +
		"PID\t0xA053\r\n" +
		"FW\t146\r\n" +
		"SER#\t12288FHJQXX\r\n" +
		"V\t13210\r\n" +
		"I\t2360\r\n" +
		"VPV\t31300\r\n" +
		"PPV\t37\r\n" +
		"CS\t3\r\n" +
		"MPPT\t2\r\n" +
		"OR\t0x00000000\r\n" +
		"ERR\t0\r\n" +
		"LOAD\tON\r\n" +
		"IL\t400\r\n" +
		"H19\t290\r\n" +
		"H20\t20\r\n" +
		"H21\t46\r\n" +
		"H22\t0\r\n" +
		"H23\t0\r\n" +
		"HSDS\t42\r\n" +
		"Checksum\tT\r\n")

}
