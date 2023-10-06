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
     * Complete the 'alternate' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts STRING s as parameter.
     */

    public static int alternate(String s) {
    // Write your code here
        String[] arrS = s.split("");
        Set<String> setS = new HashSet<String>(Arrays.asList(arrS)); 
        String[] arrFromSet = setS.toArray(new String[0]);
        
        int longestLength = 0;
        for(int i=0;i<arrFromSet.length-1;i++){
            in:
            for (int j=i+1;j<arrFromSet.length;j++){
                List<String> listS = new ArrayList<String>(Arrays.asList(arrS));
                Iterator<String> it = listS.iterator();

                while(it.hasNext()){
                    String str = it.next();
                    if(!str.equals(arrFromSet[j]) && !str.equals(arrFromSet[i])){
                        it.remove();
                    }
                }
                
                String[] res = listS.toArray(new String[0]);
                for (int k=1;k<res.length;k++){
                    if (res[k].equals(res[k-1])){
                        continue in;
                    }
                }
                
                if (listS.size() > longestLength){
                    longestLength = listS.size();
                }
            }
        }
        return longestLength;
    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int l = Integer.parseInt(bufferedReader.readLine().trim());

        String s = bufferedReader.readLine();

        int result = Result.alternate(s);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
