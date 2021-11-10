# Trivium cipher

## Tài liệu
- https://www.ecrypt.eu.org/stream/p3ciphers/trivium/trivium_p3.pdf
- https://www.dropbox.com/sh/zcbe74d19cyc5tv/AAB65O49wKKTQrXMQpfDNsyXa/slides?dl=0&preview=03streamcipher.pdf

## Mô tả
- Trivium là một thuật toán sinh khóa
- Các bit được sinh ra từ trivium sẽ được dùng để  làm khóa XOR với bản rõ

## Cài đặt
- Sử dụng 1 mảng uint64(64 bit) để lưu trữ các bit của trivium
- Khởi tạo
  - Load 80 bit của key vào vị trí 1 -> 80
  - Load 80 bit của initial value vào vị trí -> 94 - 177
  - Đặt 3 bit cuối 287 - 288 về giá trị 1
  - Warm up bằng cách chạy trivium 4*288 lần
  - ```
    // load key
    (s1, s2, . . . , s93)       ←   (K1, . . . , K80, 0, . . . , 0)
    // load initial value
    (s94, s95, . . . , s177)    ←   (IV1, . . . , IV80, 0, . . . , 0)
    // set 3 last bits
    (s178, s279, . . . , s288)  ←   (0, . . . , 0, 1, 1, 1)
    // warmup
    for i = 1 to 4 · 288 do
      nextbit
    ```
- Sinh bit
  - ```
    // generate t
    t1 ← s66 + s93
    t2 ← s162 + s177
    t3 ← s243 + s288

    // return value
    z  ← t1 + t2 + t3

    // update trivium
    t1 ← t1 + s91 · s92 + s171
    t2 ← t2 + s175 · s176 + s264
    t3 ← t3 + s286 · s287 + s69
    (s1, s2, . . . , s93) ← (t3, s1, . . . , s92)
    (s94, s95, . . . , s177) ← (t1, s94, . . . , s176)
    (s178, s279, . . . , s288) ← (t2, s178, . . . , s287)
    ```

## Mã hóa và giải mã
- Mã hóa
  - Trivium nhận key và initial value có độ dài 80 bit cho trước
  - Khởi tạo trivium
  - Ghi inital value vào đầu của cipher text
  - Mã hóa theo từng byte của clear text XOR với từng byte mà trivium sinh ra
  - Encode hex và ghi lại vào cipher text

- Giải mã
  - Trivium nhận key 80 bit cho trước
  - Đọc initial value ở đầu của cipher text
  - Decode hex của IV
  - Khởi tạo trivium
  - Đọc cipher text, decode hex và XOR tuần tự với các giá trị trivium sinh ra

## Nguồn tham khảo
- https://github.com/bmkessler/trivium
- https://github.com/sinhbad/trivium/blob/master/encrypt.c
- https://www.lipsum.com/feed/html
