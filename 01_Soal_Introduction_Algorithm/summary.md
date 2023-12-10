# Intoduction to Algorithm and Golang

INTRODICTION ALGORITHM

Pengenalan Algoritma
Input -> Proses -> Output

Algorima adalah prosedure komputasi yang didefinisikan dengan baik yang mengambil beberapa nilai sebagai input dan menghasilkan beberapa nilai sebagai output.

Contoh Algoritma
- Check Prome Number
- Sorting
- Searching

Karakteristik Algoritma
- Memiliki batas (Awal dan akhir)
- Instruksi terdefinisi dengan baik
- Efektif & Efisien

Algoritma Dasar
- Sequential
- Brancing
- Looping


Pseudocode
1. Step forward                 
2. Turn Left                    x <---------- lt
3. Step forward                 |              ^
4. Turn Left                    |              |
5. Step forward                 |              |
6. Turn left                    V              |
7. Step forward                 lt ----------->lt
8. Turn left

Kasus 1 MENGHITUNG LUAS SEGITIGA
1. INPUT Alas and Tinggi
2. CALCULATE Luas = (Alas x Tinggi) / 2
3. PRINT Luas


FLowchart
FLowchart adalah suatu bagan dengan simbool tertentu yang menggambarkan urutan dan hubungan antar proses secara mendetail

Simbol Flowchart
1. Oval -> Mulai/Selesai
2. Persegi -> Proses
3. Jajargenjang -> Input/Output
4. Belah Ketupat -> Perkondisian
5. Segi 6 -> Looping

Kasus 1 MENGHITUNG LUAS SEGITIGA
        Oval -> Mulai
                  |
                  V
Jajargenjang -> Input (A, T)
                  |
                  V
     Persegi -> L = (A * T)/2
                  |
                  V
Jajargenjang -> Output (L)
                  |
                  V
        Oval -> Selesai


Kasus 2 MENENTUKAN BILANGAN GANJIL
        Oval -> Mulai
                  |
                  V
Jajargenjang -> Input (N)
                  |
                  V
Belahketupat -> N % 2 == 0  -> Yes -> Jajargenjang -> Output("Genap")
                  |                                         |
                  V                                         |
Jajargenjang -> Output (L)                                  |
                  |                                         |
                  V                                         |
        Oval -> Selesai <------------------------------------


Kasus 3 MENCETAK FAKTOR BILANGAN
        Oval -> Mulai
                  |
                  V
Jajargenjang -> Input (N)
                  |
                  V
|----> Segi 6 -> i = 1 to N 
|                 |
|                 V 
|Belahketupat -> N % i == 0 -> Yes -|
|                 |                 |
|                 No                |
|                 |                 |
|                 V                 |
|-Belahketupat-- i <= N <-----------|
                  |
                  V
            Oval -> Selesai