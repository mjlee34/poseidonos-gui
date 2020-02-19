package v1

import "github.com/gin-gonic/gin"

func ReportTest(ctx *gin.Context) {
	getReportTest(ctx, "REPORTTEST")
}

func getReportTest(ctx *gin.Context, command string) {
	iBoFRequest := makeRequest(ctx, command)
	sendIBoF(ctx, iBoFRequest)
}

func ReportTestiCLI() {
	getReportTestCLI("REPORTTEST")
}

func getReportTestCLI(command string) {
	iBoFRequest := makeRequestCLI(command)
	sendIBoFCLI(iBoFRequest)
}
