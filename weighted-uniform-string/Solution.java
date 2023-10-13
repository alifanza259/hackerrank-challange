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
     * Complete the 'weightedUniformStrings' function below.
     *
     * The function is expected to return a STRING_ARRAY.
     * The function accepts following parameters:
     *  1. STRING s
     *  2. INTEGER_ARRAY queries
     */

    public static List<String> weightedUniformStrings(String s, List<Integer> queries) {
        // Write your code here
        String alphabet = "abcdefghijklmnopqrstuvwxyz";
        String[] alphabetArr = alphabet.split("");
        List<String> alphabetList = new ArrayList<String>(Arrays.asList(alphabetArr));
        
        String[] sArr = s.split("");
        
        Set<Integer> setOfWeight = new HashSet<Integer>();
        
        String currentChar=sArr[0];
        int weight=0;
        for (int i=0;i<sArr.length;i++){
            if(currentChar.equals(sArr[i])){
                weight+=alphabetList.indexOf(sArr[i])+1;
            }else{
                currentChar=sArr[i];
                weight=alphabetList.indexOf(sArr[i])+1;                
            }
            setOfWeight.add(weight);
        }
        
        List<String> answerArr = new ArrayList<String>();
        for(int i=0;i<queries.size();i++){
            if(setOfWeight.contains(queries.get(i))){
                answerArr.add("Yes");
            }else{
                answerArr.add("No");
            }
        }
        
        return answerArr;
    }
}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String s = bufferedReader.readLine();

        int queriesCount = Integer.parseInt(bufferedReader.readLine().trim());

        List<Integer> queries = IntStream.range(0, queriesCount).mapToObj(i -> {
            try {
                return bufferedReader.readLine().replaceAll("\\s+$", "");
            } catch (IOException ex) {
                throw new RuntimeException(ex);
            }
        })
            .map(String::trim)
            .map(Integer::parseInt)
            .collect(toList());

        List<String> result = Result.weightedUniformStrings(s, queries);

        bufferedWriter.write(
            result.stream()
                .collect(joining("\n"))
            + "\n"
        );

        bufferedReader.close();
        bufferedWriter.close();
    }
}
