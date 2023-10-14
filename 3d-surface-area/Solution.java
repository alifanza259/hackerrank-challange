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
     * Complete the 'funnyString' function below.
     *
     * The function is expected to return a STRING.
     * The function accepts STRING s as parameter.
     */

    public static String funnyString(String s) {
    // Write your code here
        String alphabet = "abcdefghijklmnopqrstuvwxyz";
        String[] alphabetArr = alphabet.split("");
        List<String> listAlphabet = new ArrayList<String>(Arrays.asList(alphabetArr));
        String[] sArr = s.split("");
        for(int i=1;i<sArr.length;i++){
            int currentCharIndex = listAlphabet.indexOf(sArr[i]);
            int previousCharIndex = listAlphabet.indexOf(sArr[i-1]);
            int diff = currentCharIndex - previousCharIndex;
            
            int reversedCurrentCharIndex = listAlphabet.indexOf(sArr[sArr.length - 1 - i]); 
            int reversedPreviousCharIndex = listAlphabet.indexOf(sArr[sArr.length - 1 - (i-1)]);
            int reversedDiff = reversedCurrentCharIndex - reversedPreviousCharIndex;

            if(Math.abs(diff) != Math.abs(reversedDiff)){
                return "Not Funny";
            }
            
        }
        
        return "Funny";
    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int q = Integer.parseInt(bufferedReader.readLine().trim());

        IntStream.range(0, q).forEach(qItr -> {
            try {
                String s = bufferedReader.readLine();

                String result = Result.funnyString(s);

                bufferedWriter.write(result);
                bufferedWriter.newLine();
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }
        });

        bufferedReader.close();
        bufferedWriter.close();
    }
}
