#include<iostream>
#include<string>
// >^.^<  kitten says: AC or treat
using namespace std;
int main(){ios::sync_with_stdio(0);cin.tie(0);int F,I,T,i,j,c,r=0;cin>>F>>I>>T;string a[10];for(i=0;i<F;)cin>>a[i++];for(j=0;j<I;j++){for(c=i=0;i<F;)c+=a[i++][j]=='Y';r+=c>=T;}cout<<r<<'\n';}
