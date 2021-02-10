#include <Windows.h>
#include <stdio.h>

int spawn_process() {
	HANDLE g_hChildStd_IN_Wr;
	HANDLE g_hChildStd_IN_Rd;
	HANDLE g_hChildStd_OUT_Wr;
	HANDLE g_hChildStd_OUT_Rd;

	SECURITY_ATTRIBUTES saAttr;

	saAttr.nLength = sizeof(SECURITY_ATTRIBUTES);
	saAttr.bInheritHandle = TRUE;
	saAttr.lpSecurityDescriptor = NULL;

	if (!CreatePipe(&g_hChildStd_OUT_Rd, &g_hChildStd_OUT_Wr, &saAttr, 0))
		return 1;
	if (!SetHandleInformation(g_hChildStd_OUT_Rd, HANDLE_FLAG_INHERIT, 0))
		return 1;
	if (!CreatePipe(&g_hChildStd_IN_Rd, &g_hChildStd_IN_Wr, &saAttr, 0))
		return 1;
	if (!SetHandleInformation(g_hChildStd_IN_Wr, HANDLE_FLAG_INHERIT, 0))
		return 1;

	TCHAR cmdLine[] = TEXT("C:\\Windows\\System32\\cmd.exe /k whoami");
	PROCESS_INFORMATION pi;
	STARTUPINFO si;
	BOOL bSuccess = FALSE;

	ZeroMemory(&pi, sizeof(PROCESS_INFORMATION));

	ZeroMemory(&si, sizeof(STARTUPINFO));
	si.cb = sizeof(STARTUPINFO);
	si.hStdError = g_hChildStd_OUT_Wr;
	si.hStdOutput = g_hChildStd_OUT_Wr;
	si.hStdInput = g_hChildStd_IN_Rd;
	si.dwFlags |= STARTF_USESTDHANDLES;

	bSuccess = CreateProcess(
		NULL,
		cmdLine,
		NULL,
		NULL,
		TRUE,
		0,
		NULL,
		NULL,
		&si,
		&pi
	);

	if (!bSuccess)
		return 2;
	else
	{
		CloseHandle(pi.hProcess);
		CloseHandle(pi.hThread);

		CloseHandle(g_hChildStd_OUT_Wr);
		CloseHandle(g_hChildStd_IN_Rd);
		g_hChildStd_OUT_Wr = 0;
		g_hChildStd_IN_Rd = 0;
	}

	DWORD dwRead, dwWritten;
	CHAR chBuf[4096];
	BOOL cSuccess = FALSE;
	HANDLE hParentStdOut = GetStdHandle(STD_OUTPUT_HANDLE);

	ZeroMemory(&chBuf, sizeof(chBuf));

	for (;;)
	{
		cSuccess = ReadFile(g_hChildStd_OUT_Rd, chBuf, 4096, &dwRead, NULL);
		if (!cSuccess || dwRead == 0)
			break;
		else
			printf("\n------\nchBuf: %s\n------\n", chBuf);

		cSuccess = WriteFile(hParentStdOut, chBuf, dwRead, &dwWritten, NULL);
		if (!cSuccess)
			break;
		else
			ZeroMemory(&chBuf, sizeof(chBuf));
	}

	return 0;
}

//int main(int argc, char const* argv[])
//{
//	return spawn_process();
//}