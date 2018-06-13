package helpers

import (
	"bytes"
	"log"
	"net/smtp"
)

func Mail(userName string, materialName string) {
	//メール機能
	//Connect to the remote SMTP server.
	d, err := smtp.Dial("sapphire.u-gakugei.ac.jp:25")
	if err != nil {
		log.Fatal(err)
	}
	//Set the sender and recipient.
	d.Mail("SemiRevel@sapphire.u-gakugei.ac.jp") // メールの送り主を指定
	d.Rcpt("hazelab@sapphire.u-gakugei.ac.jp")   // 受信者を指定

	// Send the email body.
	wc, err := d.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	//ToにするかCcにするかBccにするかはDATAメッセージ次第
	buf := bytes.NewBufferString("To:hazelab@sapphire.u-gakugei.ac.jp")
	buf.WriteString("\r\n") // DATA メッセージはCRLFのみ
	buf.WriteString("\r\n")
	buf.WriteString("Subject:" + "ゼミ資料管理システム") //件名
	buf.WriteString("\r\n")
	buf.WriteString(userName + "さんが新しい資料(" + materialName + ")を登録しました\n")
	buf.WriteString("http://onyx.u-gakugei.ac.jp/SemiRevel/ からご確認ください\n")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}

	d.Quit() //メールセッションの終了
}
