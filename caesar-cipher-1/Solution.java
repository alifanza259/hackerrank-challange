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
    final static String alphabet = "abcdefghijklmnopqrstuvwxyz";
    /*
     * Complete the 'caesarCipher' function below.
     *
     * The function is expected to return a STRING.
     * The function accepts following parameters:
     *  1. STRING s
     *  2. INTEGER k
     */

    public static String caesarCipher(String s, int k) {
    // Write your code here
        String[] alphabetArr = alphabet.split("");
        List<String> listAlphabet = new ArrayList<>(Arrays.asList(alphabetArr));
        
        String[] sArr = s.split("");
        
        for(int i=0;i<sArr.length;i++){
            int indexChar = listAlphabet.indexOf(sArr[i].toLowerCase());
            if(indexChar == -1){
                continue;
            }
            
            
            int newIndexChar = (indexChar+k)%26;
            if(sArr[i].equals(alphabetArr[indexChar])){
                sArr[i] = alphabetArr[newIndexChar];                
            }else{
                sArr[i] = alphabetArr[newIndexChar].toUpperCase();                
            }
        }
        
        return String.join("",sArr);
    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int n = Integer.parseInt(bufferedReader.readLine().trim());

        String s = bufferedReader.readLine();

        int k = Integer.parseInt(bufferedReader.readLine().trim());

        String result = Result.caesarCipher(s, k);

        bufferedWriter.write(result);
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
