## Hacking many time pad
https://github.com/thanhpp/HUST/tree/main/MMVUD/project1
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
- Từ đó, ta có thể dựa vào vị trí các ký tự \[a-z\], \[A-Z\] sau khi XOR 2 bản target với các ciphertext để  dự đoán được các ký tự có trong target

### Cách tiếp cận 1:
- [File](approach1/main.go)
- Lấy target XOR với từng cipherText, từ kết quả
  - Lưu lại vị trí của các ký tự A-Z tìm thấy
  - Cũng như số thứ tự của cipherText đang được XOR
- Từ các vị trí của ký tự A-Z ta giả định rằng
  - Tại một vị trí i, thấy được số lần xuất hiện của ký tự lớn hơn n. Thì có thể cho rằng, ở vị trí đó của cleartext có chứa dấu cách
  - Tại một vị trí j, nếu số lần xuất hiện của 1 ký tự nhó hơn m. Thì ta đoán đó là 1 ký tự của cleartext được XOR với dấu cách trong bản cleartext


### Cách tiếp cập 2:
- [File](approach2/main.go)
- Theo cách 1, ta thấy rằng có thể tìm được vị trí của dấu cách trong cleartext khi lấy ciphertext XOR với nhau
- Từ đó, ta có thể dùng phương pháp này để tìm được hết các vị trí của dấu cách trong mỗi ciphertext
- Mặt khác, ta có
```
' ' ^ key = cipher
=> cipher ^ ' ' = key
```
- Dựa vào toàn bộ vị trí các dấu cách của từng cipher text, ta có thể XOR ciphertext với ' ' để ra được key

### Tài liệu tham khảo
- https://www.fatalerrors.org/
- https://crypto.stackexchange.com/
- https://gist.github.com/python273/a326635b04ff9a80f6614a3d5bd3a840