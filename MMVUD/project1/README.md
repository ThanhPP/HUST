## Hacking many time pad

### Vấn đề:
- One time pad sử dụng lại 1 key nhiều lần để sinh ra ciphertext từ cleartext

### Ý tưởng:
- Thuật toán sinh ciphertext:
    ```
        ClearText XOR Key = CipherText
    ```
- Khi sử dụng lại 1 key nhiều lần ta có:
    ```
        ClearText1 XOR Key = CipherText1
        ClearText2 XOR Key = CipherText2
        CipherText1 XOR CipherText2 = ClearText1 XOR Key XOR ClearText2 XOR Key
                                    = ClearText1 XOR ClearText2
    ```
- Mặt khác, ta cũng có
    ```
    'a' XOR ' ' = 'A'
    'A' XOR ' ' = 'a'
    ```
- Mà ta biết cleartext chỉ bao gồm các ký tự a-zA-z
- Từ đó, ta có thể dựa vào vị trí các ký tự \[a-z\], \[A-Z\] sau khi XOR 2 target với các ciphertext để  dự đoán được các ký tự có trong target

### Cách tiếp cận 1:
- [File](main.go)
- Lấy target XOR với từng cipherText, từ kết quả
  - Lưu lại vị trí của các ký tự A-Z tìm thấy
  - Cũng như số thứ tự của cipherText đang được XOR
- Từ các vị trí của ký tự A-Z ta giả định rằng
  - Tại một vị trí i, thấy được số lần xuất hiện của ký tự lớn hơn n. Thì có thể cho rằng, ở vị trí đó của cleartext có chứa dấu cách
  - Tại một vị trí j, nếu số lần xuất hiện của 1 ký tự nhó hơn m. Thì ta đoán đó là 1 ký tự của cleartext được XOR với dấu cách trong bản cleartext

