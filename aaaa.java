import java.util.*;
class aaaa{
    public static boolean check(String s){
        char ch[]=s.toCharArray();
        int i=0,j=ch.length-1;
        while(i<=j){
            if(ch[i]!=ch[j]){
                return false;
            }
            j--;
            i++;

        }
        return true;
    }

    public static void main(String[] args){
        Scanner sc=new Scanner(System.in);
        String s=sc.next();
        if(check(s)){
            System.out.println("1");
        }
        else{
            System.out.println(0);
        }
        
    }
}