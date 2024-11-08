#include <iostream>
#include <cmath>
#include <iomanip>

using namespace std;

int main()
{
    int i, nMax=20, N=3;
    double a[N+1],
           x[nMax],
           ER[nMax],
           b[nMax][N+1],
           c[nMax][N+1],
           eps=0.003;

    x[0] = 0.5;

    a[0] = 4.0;
    a[1] = -4.0;
    a[2] = -2.0;
    a[3] = 1.0;

    b[0][N] = a[N];
    for(i=N-1; i>=0; i--){
        b[0][i] = a[i] + b[0][i+1] * x[0];
    }

    c[0][N] = a[N];
     for(i=N-1; i>0; i--){
        c[0][i] = b[0][i] + c[0][i+1] * x[0];
    }

    ER[0] = 1.0;
    i = 0;
    while(ER[i] > eps && i<nMax){
        x[i+1] = x[i] - b[i][0]/c[i][1];
        i++;

        ER[i] = abs(x[i]-x[i-1])/abs(x[i]);

        b[i][N] = a[N];
        for(int j=N-1; j>=0; j--){
            b[i][j] = a[j] + b[i][j+1] * x[i];
        }

        c[i][N] = a[N];
         for(int j=N-1; j>0; j--){
            c[i][j] = b[i][j] + c[i][j+1] * x[i];
        }
    }

    cout <<  "i\tb3\tb2\tb1\tb0" << endl;
    for(int j=0; j<=i; j++){
        cout << j << "\t";
        for(int k=N; k>=0; k--){
            cout << setprecision(10) << b[j][k] << "\t";
        }
        cout << endl;
    }

    cout << "-----------" << endl;
    cout <<  "i\tc3\tc2\tc1" << endl;
    for(int j=0; j<=i; j++){
        cout << j << "\t";
        for(int k=N; k>0; k--){
            cout << setprecision(10) << c[j][k] << "\t";
        }
        cout << endl;
    }

    return 0;
}