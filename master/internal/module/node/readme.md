endpointler:
                //her node belli aralıklarla (env da tanımlı) hearthbeat atmalı
                //bu heartheat belli süre içerisinde gelmzse (env da tanımlı)
                //node ölü kabul edilecek
                //içerisinde node-id bulunacak
                //içerisinde available space bulunacak
    /hearthbeat

                //her node ilk bağlanmada mastera ben buradayım şeklinde bir
                //register göndermeli
                //node_id içermeli
                //chunk_id listesi
                //available space içermeli
    /register

                //her node aldığı dosyanın dosya ismi ve chunkid sini gönderip
                //bu dosyayı aldığını kabul eder ve masterın belirlediği tüm
                //node ve chunklar replication factor kadar ack edilmezse hata
                //kabul edilir
    /ack
