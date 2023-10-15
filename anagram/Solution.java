package anagram;

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
     * Complete the 'anagram' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts STRING s as parameter.
     */

    public static int anagram(String s) {
    // Write your code here
        String[] sArr = s.split("");
        if(sArr.length % 2 == 1){
            return -1;
        }
        List<String> sList = new ArrayList<String>(Arrays.asList(sArr));
        
        Map<String, Integer> firstPartMap = new HashMap<String,Integer>();
        Map<String, Integer> secondPartMap = new HashMap<String,Integer>();
        for(int i=0;i<sList.size();i++){
            if(i<=sList.size()/2-1){
                firstPartMap.putIfAbsent(sList.get(i),0);
                firstPartMap.compute(sList.get(i),(k,v)->v=v+1);
            }else{
                secondPartMap.putIfAbsent(sList.get(i),0);
                secondPartMap.compute(sList.get(i),(k,v)->v=v+1);
            }
        }
        
        final int[] counter = {0}; // Declare as an array to make it effectively final

        firstPartMap.forEach((k,v)->{
            if(secondPartMap.containsKey(k)){
                if(v>secondPartMap.get(k)){
                    counter[0] += v-secondPartMap.get(k);
                }
            }else{
                counter[0] += v;
            }
        });
        
     
        return counter[0];
    }

}

class SortLexicographically implements Comparator<String> {
    public int compare(String a, String b){
        return a.compareTo(b);
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

                int result = Result.anagram(s);

                bufferedWriter.write(String.valueOf(result));
                bufferedWriter.newLine();
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }
        });

        bufferedReader.close();
        bufferedWriter.close();
    }
}
