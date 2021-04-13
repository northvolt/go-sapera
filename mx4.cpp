#include "sap.h"
#include "mx4.h"

MX4MetadataWrapper GetMX4MetadataFromBuffer(SapBufferWrapper buf, int width, int height) {
	int numRead;
	MX4MetadataWrapper metadata = new MX4Metadata;
	buf->ReadLine(width, height-1, width+64-1, height-1, metadata, &numRead);
	return metadata;
}

void MX4Metadata_Close(MX4MetadataWrapper mta) {
	delete mta;
}
