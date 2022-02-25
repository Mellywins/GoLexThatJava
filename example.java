class Factorial{
    public static void main(String[] a) {
        System.out.println(new Fac().ComputeFac(10)); 
    }
}
class Thread {
    int randomVariable;
    public void run(){
        System.out.println(randomVariable);
    }
}
class Fac extends Thread {
    public int ComputeFac(int num){
        int num_aux ;
        // base case
        
        for(int i=1; i<=num; i++){
            num_aux=i;
            System.out.println(num_aux);
        }
        if (num <= 1)
            num_aux = 1 ;
            /* else case
            
            */

        else
            num_aux = num * (this.ComputeFac(num- 1)) ;
        return num_aux ;
    }
}