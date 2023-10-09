import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.function.*;
import java.util.regex.*;
import java.util.stream.*;
import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toList;

class Result {

    /*
     * Complete the 'marsExploration' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts STRING s as parameter.
     */

    
    public static int marsExploration(String s) {
        // Write your code here
        String[] sArr = s.split("");
        int numOfLettersChanged = 0;
        
        for(int i=0;i<sArr.length;i++){
            int cursor = i % 3;
            
            if(cursor % 2 == 0 ){
                if(!sArr[i].equals("S")){
                    numOfLettersChanged++;
                }
            }else{
                if(!sArr[i].equals("O")){
                    numOfLettersChanged++;
                }
            }            
        }
        
        return numOfLettersChanged;
    }
}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String s = bufferedReader.readLine();

        int result = Result.marsExploration(s);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
