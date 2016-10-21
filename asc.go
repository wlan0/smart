package main

var errmap map[string]string

func GetErrString(a, b byte) string {
	return errmap[stringify(a, b)]
}

func init() {
	errmap = make(map[string]string)
	errmap[stringify(0x00, 0x00)] = "NO ADDITIONAL SENSE INFORMATION"
	errmap[stringify(0x00, 0x01)] = "FILEMARK DETECTED"
	errmap[stringify(0x00, 0x02)] = "END-OF-PARTITION/MEDIUM DETECTED"
	errmap[stringify(0x00, 0x03)] = "SETMARK DETECTED"
	errmap[stringify(0x00, 0x04)] = "BEGINNING-OF-PARTITION/MEDIUM DETECTED"
	errmap[stringify(0x00, 0x05)] = "END-OF-DATA DETECTED"
	errmap[stringify(0x00, 0x06)] = "I/O PROCESS TERMINATED"
	errmap[stringify(0x00, 0x11)] = "AUDIO PLAY OPERATION IN PROGRESS"
	errmap[stringify(0x00, 0x12)] = "AUDIO PLAY OPERATION PAUSED"
	errmap[stringify(0x00, 0x13)] = "AUDIO PLAY OPERATION SUCCESSFULLY COMPLETED"
	errmap[stringify(0x00, 0x14)] = "AUDIO PLAY OPERATION STOPPED DUE TO ERROR"
	errmap[stringify(0x00, 0x15)] = "NO CURRENT AUDIO STATUS TO RETURN"
	errmap[stringify(0x01, 0x00)] = "NO INDEX/SECTOR SIGNAL"
	errmap[stringify(0x02, 0x00)] = "NO SEEK COMPLETE"
	errmap[stringify(0x03, 0x00)] = "PERIPHERAL DEVICE WRITE FAULT"
	errmap[stringify(0x03, 0x01)] = "NO WRITE CURRENT"
	errmap[stringify(0x03, 0x02)] = "EXCESSIVE WRITE ERRORS"
	errmap[stringify(0x04, 0x00)] = "LOGICAL UNIT NOT READY, CAUSE NOT REPORTABLE"
	errmap[stringify(0x04, 0x01)] = "LOGICAL UNIT IS IN PROCESS OF BECOMING READY"
	errmap[stringify(0x04, 0x02)] = "LOGICAL UNIT NOT READY, INITIALIZING COMMAND REQUIRED"
	errmap[stringify(0x04, 0x03)] = "LOGICAL UNIT NOT READY, MANUAL INTERVENTION REQUIRED"
	errmap[stringify(0x04, 0x04)] = "LOGICAL UNIT NOT READY, FORMAT IN PROGRESS"
	errmap[stringify(0x05, 0x00)] = "LOGICAL UNIT DOES NOT RESPOND TO SELECTION"
	errmap[stringify(0x06, 0x00)] = "REFERENCE POSITION FOUND"
	errmap[stringify(0x07, 0x00)] = "MULTIPLE PERIPHERAL DEVICES SELECTED"
	errmap[stringify(0x08, 0x00)] = "LOGICAL UNIT COMMUNICATION FAILURE"
	errmap[stringify(0x08, 0x01)] = "LOGICAL UNIT COMMUNICATION TIME-OUT"
	errmap[stringify(0x08, 0x02)] = "LOGICAL UNIT COMMUNICATION PARITY ERROR"
	errmap[stringify(0x09, 0x00)] = "TRACK FOLLOWING ERROR"
	errmap[stringify(0x09, 0x01)] = "TRA CKING SERVO FAILURE"
	errmap[stringify(0x09, 0x02)] = "FOC US SERVO FAILURE"
	errmap[stringify(0x09, 0x03)] = "SPI NDLE SERVO FAILURE"
	errmap[stringify(0x0A, 0x00)] = "ERROR LOG OVERFLOW"
	errmap[stringify(0x0B, 0x00)] = ""
	errmap[stringify(0x0C, 0x00)] = "WRITE ERROR"
	errmap[stringify(0x0C, 0x01)] = "WRITE ERROR RECOVERED WITH AUTO REALLOCATION"
	errmap[stringify(0x0C, 0x02)] = "WRITE ERROR - AUTO REALLOCATION FAILED"
	errmap[stringify(0x0D, 0x00)] = ""
	errmap[stringify(0x0E, 0x00)] = ""
	errmap[stringify(0x0F, 0x00)] = ""
	errmap[stringify(0x10, 0x00)] = "ID CRC OR ECC ERROR"
	errmap[stringify(0x11, 0x00)] = "UNRECOVERED READ ERROR"
	errmap[stringify(0x11, 0x01)] = "READ RETRIES EXHAUSTED"
	errmap[stringify(0x11, 0x02)] = "ERROR TOO LONG TO CORRECT"
	errmap[stringify(0x11, 0x03)] = "MULTIPLE READ ERRORS"
	errmap[stringify(0x11, 0x04)] = "UNRECOVERED READ ERROR - AUTO REALLOCATE FAILED"
	errmap[stringify(0x11, 0x05)] = "L-EC UNCORRECTABLE ERROR"
	errmap[stringify(0x11, 0x06)] = "CIRC UNRECOVERED ERROR"
	errmap[stringify(0x11, 0x07)] = "DATA RESYCHRONIZATION ERROR"
	errmap[stringify(0x11, 0x08)] = "INCOMPLETE BLOCK READ"
	errmap[stringify(0x11, 0x09)] = "NO GAP FOUND"
	errmap[stringify(0x11, 0x0A)] = "MISCORRECTED ERROR"
	errmap[stringify(0x11, 0x0B)] = "UNRECOVERED READ ERROR - RECOMMEND REASSIGNMENT"
	errmap[stringify(0x11, 0x0C)] = "UNRECOVERED READ ERROR - RECOMMEND REWRITE THE DATA"
	errmap[stringify(0x12, 0x00)] = "ADDRESS MARK NOT FOUND FOR ID FIELD"
	errmap[stringify(0x13, 0x00)] = "ADDRESS MARK NOT FOUND FOR DATA FIELD"
	errmap[stringify(0x14, 0x00)] = "RECORDED ENTITY NOT FOUND"
	errmap[stringify(0x14, 0x01)] = "RECORD NOT FOUND"
	errmap[stringify(0x14, 0x02)] = "FILEMARK OR SETMARK NOT FOUND"
	errmap[stringify(0x14, 0x03)] = "END-OF-DATA NOT FOUND"
	errmap[stringify(0x14, 0x04)] = "BLOCK SEQUENCE ERROR"
	errmap[stringify(0x15, 0x00)] = "RANDOM POSITIONING ERROR"
	errmap[stringify(0x15, 0x01)] = "MECHANICAL POSITIONING ERROR"
	errmap[stringify(0x15, 0x02)] = "POSITIONING ERROR DETECTED BY READ OF MEDIUM"
	errmap[stringify(0x16, 0x00)] = "DATA SYNCHRONIZATION MARK ERROR"
	errmap[stringify(0x17, 0x00)] = "RECOVERED DATA WITH NO ERROR CORRECTION APPLIED"
	errmap[stringify(0x17, 0x01)] = "RECOVERED DATA WITH RETRIES"
	errmap[stringify(0x17, 0x02)] = "RECOVERED DATA WITH POSITIVE HEAD OFFSET"
	errmap[stringify(0x17, 0x03)] = "RECOVERED DATA WITH NEGATIVE HEAD OFFSET"
	errmap[stringify(0x17, 0x04)] = "RECOVERED DATA WITH RETRIES AND/OR CIRC APPLIED"
	errmap[stringify(0x17, 0x05)] = "RECOVERED DATA USING PREVIOUS SECTOR ID"
	errmap[stringify(0x17, 0x06)] = "RECOVERED DATA WITHOUT ECC - DATA AUTO-REALLOCATED"
	errmap[stringify(0x17, 0x07)] = "RECOVERED DATA WITHOUT ECC - RECOMMEND REASSIGNMENT"
	errmap[stringify(0x17, 0x08)] = "RECOVERED DATA WITHOUT ECC - RECOMMEND REWRITE"
	errmap[stringify(0x18, 0x00)] = "RECOVERED DATA WITH ERROR CORRECTION APPLIED"
	errmap[stringify(0x18, 0x01)] = "RECOVERED DATA WITH ERROR CORRECTION & RETRIES APPLIED"
	errmap[stringify(0x18, 0x02)] = "RECOVERED DATA - DATA AUTO-REALLOCATED"
	errmap[stringify(0x18, 0x03)] = "RECOVERED DATA WITH CIRC"
	errmap[stringify(0x18, 0x04)] = "RECOVERED DATA WITH LEC"
	errmap[stringify(0x18, 0x05)] = "RECOVERED DATA - RECOMMEND REASSIGNMENT"
	errmap[stringify(0x18, 0x06)] = "RECOVERED DATA - RECOMMEND REWRITE"
	errmap[stringify(0x19, 0x00)] = "DEFECT LIST ERROR"
	errmap[stringify(0x19, 0x01)] = "DEFECT LIST NOT AVAILABLE"
	errmap[stringify(0x19, 0x02)] = "DEFECT LIST ERROR IN PRIMARY LIST"
	errmap[stringify(0x19, 0x03)] = "DEFECT LIST ERROR IN GROWN LIST"
	errmap[stringify(0x1A, 0x00)] = "PARAMETER LIST LENGTH ERROR"
	errmap[stringify(0x1B, 0x00)] = "SYNCHRONOUS DATA TRANSFER ERROR"
	errmap[stringify(0x1C, 0x00)] = "DEFECT LIST NOT FOUND"
	errmap[stringify(0x1C, 0x01)] = "PRIMARY DEFECT LIST NOT FOUND"
	errmap[stringify(0x1C, 0x02)] = "GROWN DEFECT LIST NOT FOUND"
	errmap[stringify(0x1D, 0x00)] = "MISCOMPARE DURING VERIFY OPERATION"
	errmap[stringify(0x1E, 0x00)] = "RECOVERED ID WITH ECC"
	errmap[stringify(0x1F, 0x00)] = ""
	errmap[stringify(0x20, 0x00)] = "INVALID COMMAND OPERATION CODE"
	errmap[stringify(0x21, 0x00)] = "LOGICAL BLOCK ADDRESS OUT OF RANGE"
	errmap[stringify(0x21, 0x01)] = "INVALID ELEMENT ADDRESS"
	errmap[stringify(0x22, 0x00)] = "ILLEGAL FUNCTION (SHOULD USE 20 00, 24 00, OR 26 00)"
	errmap[stringify(0x23, 0x00)] = ""
	errmap[stringify(0x24, 0x00)] = "INVALID FIELD IN CDB"
	errmap[stringify(0x25, 0x00)] = "LOGICAL UNIT NOT SUPPORTED"
	errmap[stringify(0x26, 0x00)] = "INVALID FIELD IN PARAMETER LIST"
	errmap[stringify(0x26, 0x01)] = "PARAMETER NOT SUPPORTED"
	errmap[stringify(0x26, 0x02)] = "PARAMETER VALUE INVALID"
	errmap[stringify(0x26, 0x03)] = "THRESHOLD PARAMETERS NOT SUPPORTED"
	errmap[stringify(0x27, 0x00)] = "WRITE PROTECTED"
	errmap[stringify(0x28, 0x00)] = "NOT READY TO READY TRANSITION(MEDIUM MAY HAVE CHANGED)"
	errmap[stringify(0x28, 0x01)] = "IMPORT OR EXPORT ELEMENT ACCESSED"
	errmap[stringify(0x29, 0x00)] = "POWER ON, RESET, OR BUS DEVICE RESET OCCURRED"
	errmap[stringify(0x2A, 0x00)] = "PARAMETERS CHANGED"
	errmap[stringify(0x2A, 0x01)] = "MODE PARAMETERS CHANGED"
	errmap[stringify(0x2A, 0x02)] = "LOG PARAMETERS CHANGED"
	errmap[stringify(0x2B, 0x00)] = "COPY CANNOT EXECUTE SINCE HOST CANNOT DISCONNECT"
	errmap[stringify(0x2C, 0x00)] = "COMMAND SEQUENCE ERROR"
	errmap[stringify(0x2C, 0x01)] = "TOO MANY WINDOWS SPECIFIED"
	errmap[stringify(0x2C, 0x02)] = "INVALID COMBINATION OF WINDOWS SPECIFIED"
	errmap[stringify(0x2D, 0x00)] = "OVERWRITE ERROR ON UPDATE IN PLACE"
	errmap[stringify(0x2E, 0x00)] = ""
	errmap[stringify(0x2F, 0x00)] = "COMMANDS CLEARED BY ANOTHER INITIATOR"
	errmap[stringify(0x30, 0x00)] = "INCOMPATIBLE MEDIUM INSTALLED"
	errmap[stringify(0x30, 0x01)] = "CANNOT READ MEDIUM - UNKNOWN FORMAT"
	errmap[stringify(0x30, 0x02)] = "CANNOT READ MEDIUM - INCOMPATIBLE FORMAT"
	errmap[stringify(0x30, 0x03)] = "CLEANING CARTRIDGE INSTALLED"
	errmap[stringify(0x31, 0x00)] = "MEDIUM FORMAT CORRUPTED"
	errmap[stringify(0x31, 0x01)] = "FORMAT COMMAND FAILED"
	errmap[stringify(0x32, 0x00)] = "NO DEFECT SPARE LOCATION AVAILABLE"
	errmap[stringify(0x32, 0x01)] = "DEFECT LIST UPDATE FAILURE"
	errmap[stringify(0x33, 0x00)] = "TAPE LENGTH ERROR"
	errmap[stringify(0x34, 0x00)] = ""
	errmap[stringify(0x35, 0x00)] = ""
	errmap[stringify(0x36, 0x00)] = "RIBBON, INK, OR TONER FAILURE"
	errmap[stringify(0x37, 0x00)] = "ROUNDED PARAMETER"
	errmap[stringify(0x38, 0x00)] = ""
	errmap[stringify(0x39, 0x00)] = "SAVING PARAMETERS NOT SUPPORTED"
	errmap[stringify(0x3A, 0x00)] = "MEDIUM NOT PRESENT"
	errmap[stringify(0x3B, 0x00)] = "SEQUENTIAL POSITIONING ERROR"
	errmap[stringify(0x3B, 0x01)] = "TAPE POSITION ERROR AT BEGINNING-OF-MEDIUM"
	errmap[stringify(0x3B, 0x02)] = "TAPE POSITION ERROR AT END-OF-MEDIUM"
	errmap[stringify(0x3B, 0x03)] = "TAPE OR ELECTRONIC VERTICAL FORMS UNIT NOT READY"
	errmap[stringify(0x3B, 0x04)] = "SLEW FAILURE"
	errmap[stringify(0x3B, 0x05)] = "PAPER JAM"
	errmap[stringify(0x3B, 0x06)] = "FAILED TO SENSE TOP-OF-FORM"
	errmap[stringify(0x3B, 0x07)] = "FAILED TO SENSE BOTTOM-OF-FORM"
	errmap[stringify(0x3B, 0x08)] = "REPOSITION ERROR"
	errmap[stringify(0x3B, 0x09)] = "READ PAST END OF MEDIUM"
	errmap[stringify(0x3B, 0x0A)] = "READ PAST BEGINNING OF MEDIUM"
	errmap[stringify(0x3B, 0x0B)] = "POSITION PAST END OF MEDIUM"
	errmap[stringify(0x3B, 0x0C)] = "POSITION PAST BEGINNING OF MEDIUM"
	errmap[stringify(0x3B, 0x0D)] = "MEDIUM DESTINATION ELEMENT FULL"
	errmap[stringify(0x3B, 0x0E)] = "MEDIUM SOURCE ELEMENT EMPTY"
	errmap[stringify(0x3C, 0x00)] = ""
	errmap[stringify(0x3D, 0x00)] = "INVALID BITS IN IDENTIFY MESSAGE"
	errmap[stringify(0x3E, 0x00)] = "LOGICAL UNIT HAS NOT SELF-CONFIGURED YET"
	errmap[stringify(0x3F, 0x00)] = "TARGET OPERATING CONDITIONS HAVE CHANGED"
	errmap[stringify(0x3F, 0x01)] = "MICROCODE HAS BEEN CHANGED"
	errmap[stringify(0x3F, 0x02)] = "CHANGED OPERATING DEFINITION"
	errmap[stringify(0x3F, 0x03)] = "INQUIRY DATA HAS CHANGED"
	errmap[stringify(0x40, 0x00)] = "RAM FAILURE (SHOULD USE 40 NN)"
	//errmap[stringify(0x40, 0xNN)] = "DIAGNOSTIC FAILURE ON COMPONENT NN (80H-FFH)"
	errmap[stringify(0x41, 0x00)] = "DATA PATH FAILURE (SHOULD USE 40 NN)"
	errmap[stringify(0x42, 0x00)] = "POWER-ON OR SELF-TEST FAILURE (SHOULD USE 40 NN)"
	errmap[stringify(0x43, 0x00)] = "MESSAGE ERROR"
	errmap[stringify(0x44, 0x00)] = "INTERNAL TARGET FAILURE"
	errmap[stringify(0x45, 0x00)] = "SELECT OR RESELECT FAILURE"
	errmap[stringify(0x46, 0x00)] = "UNSUCCESSFUL SOFT RESET"
	errmap[stringify(0x47, 0x00)] = "SCSI PARITY ERROR"
	errmap[stringify(0x48, 0x00)] = "INITIATOR DETECTED ERROR MESSAGE RECEIVED"
	errmap[stringify(0x49, 0x00)] = "INVALID MESSAGE ERROR"
	errmap[stringify(0x4A, 0x00)] = "COMMAND PHASE ERROR"
	errmap[stringify(0x4B, 0x00)] = "DATA PHASE ERROR"
	errmap[stringify(0x4C, 0x00)] = "LOGICAL UNIT FAILED SELF-CONFIGURATION"
	errmap[stringify(0x4D, 0x00)] = ""
	errmap[stringify(0x4E, 0x00)] = "OVERLAPPED COMMANDS ATTEMPTED"
	errmap[stringify(0x4F, 0x00)] = ""
	errmap[stringify(0x50, 0x00)] = "WRITE APPEND ERROR"
	errmap[stringify(0x50, 0x01)] = "WRITE APPEND POSITION ERROR"
	errmap[stringify(0x50, 0x02)] = "POSITION ERROR RELATED TO TIMING"
	errmap[stringify(0x51, 0x00)] = "ERASE FAILURE"
	errmap[stringify(0x52, 0x00)] = "CARTRIDGE FAULT"
	errmap[stringify(0x53, 0x00)] = "MEDIA LOAD OR EJECT FAILED"
	errmap[stringify(0x53, 0x01)] = "UNLOAD TAPE FAILURE"
	errmap[stringify(0x53, 0x02)] = "MEDIUM REMOVAL PREVENTED"
	errmap[stringify(0x54, 0x00)] = "SCSI TO HOST SYSTEM INTERFACE FAILURE"
	errmap[stringify(0x55, 0x00)] = "SYSTEM RESOURCE FAILURE"
	errmap[stringify(0x56, 0x00)] = ""
	errmap[stringify(0x57, 0x00)] = "UNABLE TO RECOVER TABLE-OF-CONTENTS"
	errmap[stringify(0x58, 0x00)] = "GENERATION DOES NOT EXIST"
	errmap[stringify(0x59, 0x00)] = "UPDATED BLOCK READ"
	errmap[stringify(0x5A, 0x00)] = "OPERATOR REQUEST OR STATE CHANGE INPUT (UNSPECIFIED)"
	errmap[stringify(0x5A, 0x01)] = "OPERATOR MEDIUM REMOVAL REQUEST"
	errmap[stringify(0x5A, 0x02)] = "OPERATOR SELECTED WRITE PROTECT"
	errmap[stringify(0x5A, 0x03)] = "OPERATOR SELECTED WRITE PERMIT"
	errmap[stringify(0x5B, 0x00)] = "LOG EXCEPTION"
	errmap[stringify(0x5B, 0x01)] = "THRESHOLD CONDITION MET"
	errmap[stringify(0x5B, 0x02)] = "LOG COUNTER AT MAXIMUM"
	errmap[stringify(0x5B, 0x03)] = "LOG LIST CODES EXHAUSTED"
	errmap[stringify(0x5C, 0x00)] = "RPL STATUS CHANGE"
	errmap[stringify(0x5C, 0x01)] = "SPINDLES SYNCHRONIZED"
	errmap[stringify(0x5C, 0x02)] = "SPINDLES NOT SYNCHRONIZED"
	errmap[stringify(0x5D, 0x00)] = ""
	errmap[stringify(0x5E, 0x00)] = ""
	errmap[stringify(0x5F, 0x00)] = ""
	errmap[stringify(0x60, 0x00)] = "LAMP FAILURE"
	errmap[stringify(0x61, 0x00)] = "VIDEO ACQUISITION ERROR"
	errmap[stringify(0x61, 0x01)] = "UNABLE TO ACQUIRE VIDEO"
	errmap[stringify(0x61, 0x02)] = "OUT OF FOCUS"
	errmap[stringify(0x62, 0x00)] = "SCAN HEAD POSITIONING ERROR"
	errmap[stringify(0x63, 0x00)] = "END OF USER AREA ENCOUNTERED ON THIS TRACK"
	errmap[stringify(0x64, 0x00)] = "ILLEGAL MODE FOR THIS TRACK"
}
