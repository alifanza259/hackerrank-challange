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
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Result {

    /*
     * Complete the 'minimumNumber' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts following parameters:
     * 1. INTEGER n
     * 2. STRING password
     */
    public static boolean is6Length = false;
    public static boolean containsDigit = false;
    public static boolean containsLower = false;
    public static boolean containsUpper = false;
    public static boolean containsSpecial = false;

    public static int minimumNumber(int n, String password) {
        // Return the minimum number of characters to make the password strong
        String numbers = "0123456789";
        String lower_case = "abcdefghijklmnopqrstuvwxyz";
        String upper_case = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
        String special_characters = "!@#$%^&*()-+";

        List<String> listNum = new ArrayList<String>(Arrays.asList(numbers.split("")));
        List<String> listLower = new ArrayList<String>(Arrays.asList(lower_case.split("")));
        List<String> listUpper = new ArrayList<String>(Arrays.asList(upper_case.split("")));
        List<String> listSpecial = new ArrayList<String>(Arrays.asList(special_characters.split("")));

        String[] arrPassword = password.split("");
        for (int i = 0; i < n; i++) {
            String c = arrPassword[i];
            if (!containsDigit && listNum.contains(c)) {
                containsDigit = true;
            }

            if (!containsLower && listLower.contains(c)) {
                containsLower = true;
            }

            if (!containsUpper && listUpper.contains(c)) {
                containsUpper = true;
            }

            if (!containsSpecial && listSpecial.contains(c)) {
                containsSpecial = true;
            }
        }

        int neededAddition = 0;
        if (!containsDigit) {
            neededAddition++;
        }
        if (!containsLower) {
            neededAddition++;
        }
        if (!containsUpper) {
            neededAddition++;
        }
        if (!containsSpecial) {
            neededAddition++;
        }
        if (n < 6) {
            if (neededAddition < (6 - n)) {
                neededAddition = 6 - n;
            }
        }

        return neededAddition;
    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int n = Integer.parseInt(bufferedReader.readLine().trim());

        String password = bufferedReader.readLine();

        int answer = Result.minimumNumber(n, password);

        bufferedWriter.write(String.valueOf(answer));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
