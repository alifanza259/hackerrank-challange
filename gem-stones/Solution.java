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
import java.util.concurrent.atomic.AtomicInteger;

class Result {

    /*
     * Complete the 'gemstones' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts STRING_ARRAY arr as parameter.
     */

    public static int gemstones(List<String> arr) {
    // Write your code here 
        Map<String, Integer> mineralMap = new HashMap<String, Integer>();
        
        for(int i=0;i<arr.size();i++){
            String[] mineralArr = arr.get(i).split("");
            Set<String> mineralSet = new HashSet<String>();
            for(int j=0;j<mineralArr.length;j++){
                mineralSet.add(mineralArr[j]);
            }
            
            String[] uniqueMineralArr = mineralSet.toArray(new String[0]);
            for(int j=0;j<uniqueMineralArr.length;j++){
                mineralMap.compute(uniqueMineralArr[j], (k, v) -> (v == null) ? 1 : v + 1);
            }
        }
        
        AtomicInteger counter = new AtomicInteger(0);
        mineralMap.forEach((key,val)-> {
            if(val == arr.size()){
                counter.incrementAndGet();
            }
        });
        
        return counter.get();

    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int n = Integer.parseInt(bufferedReader.readLine().trim());

        List<String> arr = IntStream.range(0, n).mapToObj(i -> {
            try {
                return bufferedReader.readLine();
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }
        })
            .collect(toList());

        int result = Result.gemstones(arr);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
    