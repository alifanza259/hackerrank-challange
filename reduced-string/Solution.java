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
     * Complete the 'superReducedString' function below.
     *
     * The function is expected to return a STRING.
     * The function accepts STRING s as parameter.
     */

    public static String superReducedString(String s) {
        // Write your code here
        String[] arrS = s.split("");
        List<String> listS = new ArrayList<String>(Arrays.asList(arrS)); 
        int i=1;
        while(true){
            if (listS.size() == 0 || i>=listS.size()){
                break;
            }
            if (listS.get(i).equals(listS.get(i-1))){
                listS.remove(i);
                listS.remove(i-1);
                if(i!=1){
                    i--;
                }
            }else{
                i++;
            }
        }
        
        if(listS.size() == 0){
            return "Empty String";
        }else{
            return String.join("", listS);
        }
    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String s = bufferedReader.readLine();

        String result = Result.superReducedString(s);

        bufferedWriter.write(result);
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
