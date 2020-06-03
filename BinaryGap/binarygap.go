package Binarygap

func BinaryGap(N int) int {
	nStart := -1 // 1 시작 지점
	nMaxlen := 0 // 현재 최대 길이

	for nPos := 0; N > 0; nPos++ {
		if N&1 == 1 {

			// 최대 길이 갱신
			if nStart != -1 {
				nLen := nPos - nStart
				if nLen > nMaxlen {
					nMaxlen = nLen
				}
			}

			nStart = nPos // 이전 1위치 갱신
		}

		N = N >> 1
	}

	return nMaxlen
}
