package Parcer;

/*
<Type>::= <Base> <Tail>
<Tail>::= -> <Type> | ε
<Base>::= <Cort> <List>
<List>::= list <List> | ε
<Cort>::= <Factor> <CTail>
<CTail>::= * <Cort> | ε
<Factor>::= int | real | ( <Type> )
*/

public class Parcer {
    private static String inputLine;
    private static int pos = 0;
    private static int line = 1;
    private static int col = 1;

    public static void main(String[] args) {

        String inputData = "int * (int list) ->\n" +
                "int * (real list -> int * real) * int * (int list -> real) ->\n" +
                "real"; // предложение для анализа

        inputLine = inputData;

        try {
            parseType();

            if (pos == inputLine.length()) {
                System.out.println("Предложение соответствует грамматике");
                System.out.println("<Type>::= <Base> <Tail>\n" +
                        "<Tail>::= -> <Type> | ε\n" +
                        "<Base>::= <Cort> <List>\n" +
                        "<List>::= list <List> | ε\n" +
                        "<Cort>::= <Factor> <CTail>\n" +
                        "<CTail>::= * <Cort> | ε\n" +
                        "<Factor>::= int | real | ( <Type> )");
            } else {
                throw new SyntaxErrorException("Предложение не соответствует грамматике");
            }

        } catch (SyntaxErrorException error) {
            System.out.println(error.getMessage());
        }
    }

    private static void parseType() {
        parseBase();
        parseTail();
    }

    private static void parseTail() {
        parseWhitespaces();
        if (pos + 1 < inputLine.length() &&
                inputLine.charAt(pos) == '-' &&
                inputLine.charAt(pos+1) == '>'){
            pos += 2;
            col += 2;
            parseType();
        }
    }

    private static void parseBase() {
        parseCort();
        parseList();
    }

    private static void parseList() {
        parseWhitespaces();
        if (pos + 3 < inputLine.length() &&
                inputLine.charAt(pos) == 'l' &&
                inputLine.charAt(pos+1) == 'i' &&
                inputLine.charAt(pos+2) == 's' &&
                inputLine.charAt(pos+3) == 't'){
            pos += 4;
            col += 4;
            parseList();
        }
    }

    private static void parseCort() {
        parseFactor();
        parseCtail();
    }

    private static void parseCtail() {
        parseWhitespaces();
        if (pos < inputLine.length() && inputLine.charAt(pos) == '*'){
            pos++;
            col++;
            parseCort();
        }
    }

    private static void parseFactor() throws SyntaxErrorException{
        parseWhitespaces();
        if (pos < inputLine.length() && inputLine.charAt(pos) == '(') {
            pos++;
            col++;
            parseType();
            if (pos < inputLine.length() && inputLine.charAt(pos) == ')') {
                pos++;
                col++;
            } else {
                throw new SyntaxErrorException("Предложение не соответствует грамматике");
            }

        } else if (pos < inputLine.length() && inputLine.charAt(pos) == 'i'){
            parseInt();

        } else if (pos < inputLine.length() && inputLine.charAt(pos) == 'r'){
            parseReal();
        } else {
            throw new SyntaxErrorException("Предложение не соответствует грамматике");
        }
    }

    private static void parseInt() {
        parseWhitespaces();
        if (pos + 2 < inputLine.length() &&
                inputLine.charAt(pos) == 'i' &&
                inputLine.charAt(pos+1) == 'n' &&
                inputLine.charAt(pos+2) == 't'){
            pos += 3;
            col += 3;
        } else {
            throw new SyntaxErrorException("Предложение не соответствует грамматике");
        }
    }

    private static void parseReal() {
        parseWhitespaces();
        if (pos + 3 < inputLine.length() &&
                inputLine.charAt(pos) == 'r' &&
                inputLine.charAt(pos+1) == 'e' &&
                inputLine.charAt(pos+2) == 'a' &&
                inputLine.charAt(pos+3) == 'l'){
            pos += 4;
            col += 4;
        } else {
            throw new SyntaxErrorException("Предложение не соответствует грамматике");
        }
    }

    private static void parseWhitespaces() {
        while (pos < inputLine.length() &&
                (inputLine.charAt(pos) == ' ' || inputLine.charAt(pos) == '\n')){
            if (inputLine.charAt(pos) == ' ') {
                pos++;
                col++;
            } else {
                pos++;
                col = 1;
                line++;
            }
        }
    }

    static class SyntaxErrorException extends RuntimeException {
        SyntaxErrorException(String message) {
            super(message + " (строка " + line + ", позиция " + col + ")");
        }
    }
}