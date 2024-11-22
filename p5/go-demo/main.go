package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

// 各文字をストリーミングで送信する関数
func streamCharacters(text string, w http.ResponseWriter, delay time.Duration) {
    for _, char := range text {
        fmt.Fprintf(w, "%c", char)
        w.(http.Flusher).Flush()
        time.Sleep(delay)
    }
    fmt.Fprintf(w, "\n")
    w.(http.Flusher).Flush()
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // ストリーミングのためのヘッダー設定
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("Cache-Control", "no-cache")
        w.Header().Set("Connection", "keep-alive")

        // 文字列の準備
        messages := []string{
            "あべさんはなんでもできるまん",
            "たなかさんはせがたかいまん",
            "まつださんはつよつよまん",
        }

        // 最終的な文字列を組み立て
        finalText := ""
        for _, msg := range messages {
            if len(msg) >= 3 {
                finalText += string(msg[0]) // 1文字目
            }
        }
        finalText += "。。。"

        // まず3つのメッセージを順番に表示（各文字を0.1秒間隔で）
        for _, msg := range messages {
            streamCharacters(msg, w, 100*time.Millisecond)
            time.Sleep(500 * time.Millisecond) // メッセージ間の待ち時間
        }

        // 少し間を置いて
        time.Sleep(1 * time.Second)

        // 最終メッセージを表示
        streamCharacters(finalText, w, 200*time.Millisecond)
    })

    log.Println("Server starting on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}