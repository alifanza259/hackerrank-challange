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
     * Complete the 'gameOfThrones' function below.
     *
     * The function is expected to return a STRING.
     * The function accepts STRING s as parameter.
     */

    public static String gameOfThrones(String s) {
    // Write your code here
        String[] sArr = s.split("");

        
        Map<String, Integer> sMap = new HashMap<String,Integer>();
        for(int i=0;i<sArr.length;i++){
            sMap.putIfAbsent(sArr[i],0);
            sMap.compute(sArr[i],(k,v)->v=v+1);
        }
        
        final String[] answer = {"YES"};
        
        if (sArr.length % 2 == 0){
            sMap.forEach((k,v)->{
                if(v%2==1){
                    answer[0]="NO";
                }
            });
        }else{
            final int[] counter = {0};
            sMap.forEach((k,v)->{
                if(v%2==1){
                    if(counter[0]>0){
                        
                    answer[0]="NO";
                    }else{                        
                        counter[0]++;
                    }
                }
            });
        }
        
        return answer[0];

    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String s = bufferedReader.readLine();

        String result = Result.gameOfThrones(s);

        bufferedWriter.write(result);
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
