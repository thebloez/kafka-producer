package main

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"kafka-producer/model"
	"math/rand"
	"time"
)

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})

	if err != nil {
		panic(err)
	}

	defer producer.Close()

	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < 1; i++ {
		transaksi := generateRandomTransaction()
		transaksiJson, errParse := json.Marshal(transaksi)
		if errParse != nil {
			fmt.Printf("Failed to marshal transaction :%s", errParse.Error())
			continue
		}

		var topic = "transactions"
		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          transaksiJson,
		}, nil)
		if err != nil {
			fmt.Printf("failed to produce message : %s", err.Error())
		}
		fmt.Println("Success produce and send message " + string(transaksiJson))
		time.Sleep(2000 * time.Millisecond)
	}

	producer.Flush(15 * 1000)
}

func generateRandomTransaction() model.Transaksi {
	//transactionTypes := []string{"pembelian", "transfer", "pembayaran_tagihan"}
	transactionTypes := []string{"pembelian"}
	jenis := transactionTypes[rand.Intn(len(transactionTypes))]

	transaksi := model.Transaksi{
		IDTransaksi: "TRX" + time.Now().Format("20060102150405") + fmt.Sprintf("%03d", rand.Intn(1000)),
		Tanggal:     time.Now().Format("2006-01-02"),
		Jenis:       jenis,
		Jumlah:      rand.Intn(10_000) + 10_000,
		MataUang:    "IDR",
	}

	switch jenis {
	case "pembelian":
		transaksi.MetodePembayaran = "kartu_kredit"
		transaksi.Detail = &model.DetailPembelian{
			NamaPedagang: gofakeit.Company(),
			Kategori:     gofakeit.RandomString([]string{"elektronik", "fashion", "makanan", "minuman"}),
			Deskripsi:    "Pembelian Barang Pribadi",
		}
	case "transfer":
		transaksi.Pengirim = &model.PengirimTransfer{
			Nama:          gofakeit.Name(),
			NomorRekening: generateRandomAccountNumber(),
		}
		transaksi.Penerima = &model.PenerimaTransfer{
			Nama:          gofakeit.Name(),
			NomorRekening: generateRandomAccountNumber(),
		}
	case "pembayaran_tagihan":
		transaksi.PenyediaJasa = gofakeit.Company()
		transaksi.NomorPelanggan = gofakeit.UUID()
		transaksi.PeriodeTagihan = "2024-06"
	}

	return transaksi
}

func generateRandomAccountNumber() string {
	return fmt.Sprintf("%010d", rand.Intn(1000000000))
}
