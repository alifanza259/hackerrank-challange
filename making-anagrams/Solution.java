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
     * Complete the 'makingAnagrams' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts following parameters:
     * 1. STRING s1
     * 2. STRING s2
     */

    public static int makingAnagrams(String s1, String s2) {
        // Write your code here
        String[] s1Arr = s1.split("");
        String[] s2Arr = s2.split("");

        Map<String, Integer> s1Map = new HashMap<String, Integer>();
        Map<String, Integer> s2Map = new HashMap<String, Integer>();
        for (int i = 0; i < s1Arr.length; i++) {
            s1Map.putIfAbsent(s1Arr[i], 0);
            s1Map.compute(s1Arr[i], (k, v) -> v = v + 1);
        }

        for (int i = 0; i < s2Arr.length; i++) {
            s2Map.putIfAbsent(s2Arr[i], 0);
            s2Map.compute(s2Arr[i], (k, v) -> v = v + 1);
        }

        final int[] counter = { 0 };

        s1Map.forEach((k, v) -> {
            if (s2Map.containsKey(k)) {
                if (v > s2Map.get(k)) {
                    counter[0] += v - s2Map.get(k);
                }
            } else {
                counter[0] += v;
            }
        });

        s2Map.forEach((k, v) -> {
            if (s1Map.containsKey(k)) {
                if (v > s1Map.get(k)) {
                    counter[0] += v - s1Map.get(k);
                }
            } else {
                counter[0] += v;
            }
        });
        return counter[0];
    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String s1 = bufferedReader.readLine();

        String s2 = bufferedReader.readLine();

        int result = Result.makingAnagrams(s1, s2);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
