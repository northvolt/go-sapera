#include "sap.h"
#include "mx4.h"

MX4MetadataWrapper GetMX4MetadataFromBuffer(SapBufferWrapper buf, int width, int height) {
	int numRead = 0;
	MX4MetadataWrapper metadata = new MX4MetadataWrapper;
	buf->ReadLine(width, height-1, width+64-1, height-1, metadata, &numRead);
	int fc = metadata.frameCounter;
	int lc = metadata.lineCount;
	return metadata;
}

void MX4Metadata_Close(MX4MetadataWrapper mta) {
	delete mta;
}
