// Shell Emulator Java Project
// This code emulators a standard shell in a UNIX/Linux Distro

// Authentication: only one user for now --- three attempts
// Input Processing: several commands --- with support for command arguments

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Scanner;

public class JavaShell {
    // Environment variables - moved to class level for easier access across methods
    private static String currentDirectory = "home/kelz/";
    private static final String userName = "kelz";
    private static final HashMap<String, String> directories = new HashMap<>();

    // Initialize the file system structure
    static {
        // Add some directories
        directories.put("home/kelz/", "Documents Screenshots JavaProjects Music GitRepos");
        directories.put("home/kelz/Documents/", "resume.txt notes.md project_ideas.txt");
        directories.put("home/kelz/JavaProjects/", "JavaShell.java Calculator.java");
    }

    // Authenticate user
    public static boolean authUser() {
        System.out.println("""
                     _                    ____  _          _ _\s
                    | | __ ___   ____ _  / ___|| |__   ___| | |
                 _  | |/ _` \\ \\ / / _` | \\___ \\| '_ \\ / _ \\ | |
                | |_| | (_| |\\ V / (_| |  ___) | | | |  __/ | |
                 \\___/ \\__,_| \\_/ \\__,_| |____/|_| |_|\\___|_|_|""");

        // Create Scanner Object for auth
        Scanner authScanner = new Scanner(System.in);

        // Receive user creds
        int attempts = 3;
        int authenticated = 0;
        while (authenticated == 0) {
            System.out.print("Username: ");
            String userName = authScanner.nextLine();
            if (userName.equals("kelz")) {
                while (attempts > 0) {
                    System.out.print("Password: ");
                    String userPass = authScanner.nextLine();
                    if (userPass.equals("pass123")) {
                        authenticated = 1;
                        return true;
                    } else {
                        System.out.println("Incorrect Password --- Attempts left = " + (attempts - 1));
                        attempts--;
                        if (attempts == 0){
                            System.out.println("Too many failed attempts. Exiting...");
                            System.exit(1); // Terminates the program
                        }
                    }
                }
            } else {
                System.out.println("Username not found...");
            }
        }
        return true;
    }

    // Parse user input to separate command and arguments
    public static String[] parseCommand(String userInput) {
        // Split input by spaces, preserving quoted text as a single argument
        List<String> matchList = new ArrayList<>();

        // Simple split for now - can be enhanced to handle quotes and escape characters
        return userInput.trim().split("\\s+", 2);
    }

    // Process user input
    public static void processInput(String userInput) {
        if (userInput.trim().isEmpty()) {
            return; // Skip processing for empty input
        }

        // Parse the command and arguments
        String[] cmdParts = parseCommand(userInput);
        String command = cmdParts[0].toLowerCase(); // Command is always the first part
        String args = cmdParts.length > 1 ? cmdParts[1] : ""; // Arguments if any

        // Process command with switch statement
        switch(command) {
            //ls -- lists contents of directory
            case "ls" -> {
                if (args.isEmpty()) {
                    // List current directory
                    if (directories.containsKey(currentDirectory)) {
                        System.out.println(directories.get(currentDirectory));
                    } else {
                        System.out.println("Directory not found in simulation.");
                    }
                } else {
                    // List specified directory
                    String dirToList = resolvePath(args);
                    if (directories.containsKey(dirToList)) {
                        System.out.println(directories.get(dirToList));
                    } else {
                        System.out.println("ls: cannot access '" + args + "': No such file or directory");
                    }
                }
            }

            // cd -- change directory
            case "cd" -> {
                if (args.isEmpty()) {
                    // Default to home directory if no args
                    currentDirectory = "home/kelz/";
                    System.out.println("Changed to home directory: " + currentDirectory);
                } else {
                    // Change to specified directory
                    String newDir = resolvePath(args);
                    if (directories.containsKey(newDir)) {
                        currentDirectory = newDir;
                        System.out.println("Changed to directory: " + currentDirectory);
                    } else {
                        System.out.println("cd: " + args + ": No such file or directory");
                    }
                }
            }

            //hostname -- return hostname
            case "hostname" -> {
                System.out.println("javashell");
            }

            //arch -- return system architecture
            case "arch" -> {
                System.out.println("x86_64");
            }

            //free -- return memory info
            case "free" -> {
                System.out.println("""
                                       total        used        free      shared  buff/cache   available
                        Mem:        14201168      728216    13175756        3252      562576    13472952
                        Swap:        4194304           0     4194304""");
            }

            //date -- return date info
            case "date" -> {
                System.out.println("Mon Mar  3 19:20:09 CST 2025");
            }

            // clear -- clears screen
            case "clear" -> {
                for(int i=0;i<20;i++){
                    System.out.println();
                }
            }

            // pwd -- print current directory
            case "pwd" -> {
                System.out.println(currentDirectory);
            }

            // ping -- test network connections
            case "ping" -> {
                if (args.isEmpty()) {
                    System.out.println("Usage: ping <host>");
                } else {
                    System.out.println("PING " + args + " (" + args + ") 56(84) bytes of data.");
                    System.out.println("64 bytes from " + args + ": icmp_seq=1 ttl=114 time=13.0 ms");
                    System.out.println("64 bytes from " + args + ": icmp_seq=2 ttl=114 time=12.5 ms");
                    System.out.println("64 bytes from " + args + ": icmp_seq=3 ttl=114 time=13.3 ms");
                    System.out.println("64 bytes from " + args + ": icmp_seq=4 ttl=114 time=11.3 ms");
                }
            }

            // whoami -- prints out username
            case "whoami" -> {
                System.out.println(userName);
            }

            // ps -- prints out processes
            case "ps" -> {
                System.out.println("""
                        PID TTY          TIME CMD
                          18879 pts/3    00:00:00 bash
                          21005 pts/3    00:00:00 ps""");
            }

            // echo -- display a line of text
            case "echo" -> {
                System.out.println(args);
            }

            //lsmem -- list memory information
            case "lsmem" -> {
                System.out.println("""
                        RANGE                                  SIZE  STATE REMOVABLE  BLOCK
                        0x0000000000000000-0x00000000f7ffffff  3.9G online       yes   0-30
                        0x0000000100000000-0x0000000387ffffff 10.1G online       yes 32-112
                        
                        Memory block size:       128M
                        Total online memory:      14G
                        Total offline memory:      0B""");
            }

            // cat -- concatenate and display file contents
            case "cat" -> {
                if (args.isEmpty()) {
                    System.out.println("Usage: cat <file>");
                } else {
                    // Simple file content simulation
                    if (args.equals("resume.txt") && currentDirectory.equals("home/kelz/Documents/")) {
                        System.out.println("This is a simulated resume file content.");
                    } else if (args.equals("notes.md") && currentDirectory.equals("home/kelz/Documents/")) {
                        System.out.println("# Notes\n- Improve Java Shell project\n- Add more commands\n- Implement file system simulation");
                    } else {
                        System.out.println("cat: " + args + ": No such file or directory");
                    }
                }
            }

            // mkdir -- create directory
            case "mkdir" -> {
                if (args.isEmpty()) {
                    System.out.println("Usage: mkdir <directory>");
                } else {
                    String newDirPath = currentDirectory + args + "/";
                    if (directories.containsKey(newDirPath)) {
                        System.out.println("mkdir: cannot create directory '" + args + "': File exists");
                    } else {
                        directories.put(newDirPath, ""); // Add empty directory
                        System.out.println("Directory created: " + newDirPath);
                    }
                }
            }

            // touch -- create empty file
            case "touch" -> {
                if (args.isEmpty()) {
                    System.out.println("Usage: touch <file>");
                } else {
                    // Add file to current directory's content
                    String dirContent = directories.getOrDefault(currentDirectory, "");
                    if (!dirContent.contains(args)) {
                        directories.put(currentDirectory, dirContent + (dirContent.isEmpty() ? "" : " ") + args);
                        System.out.println("File created: " + args);
                    } else {
                        System.out.println("touch: updated timestamp for '" + args + "'");
                    }
                }
            }

            // help -- display available commands
            case "help" -> {
                System.out.println("Available commands:");
                System.out.println("  ls [directory]    - List directory contents");
                System.out.println("  cd [directory]    - Change directory");
                System.out.println("  pwd               - Print current directory");
                System.out.println("  clear             - Clear the screen");
                System.out.println("  whoami            - Display current user");
                System.out.println("  hostname          - Display system hostname");
                System.out.println("  date              - Display current date");
                System.out.println("  ps                - Display process status");
                System.out.println("  arch              - Display system architecture");
                System.out.println("  free              - Display memory usage");
                System.out.println("  lsmem             - List memory blocks");
                System.out.println("  ping <host>       - Test network connection");
                System.out.println("  echo <text>       - Display text");
                System.out.println("  cat <file>        - Display file contents");
                System.out.println("  mkdir <directory> - Create directory");
                System.out.println("  touch <file>      - Create empty file");
                System.out.println("  exit              - Exit the shell");
            }

            // default to handle edge case
            default -> System.out.println(command + ": command not found");
        }
    }

    // Helper method to resolve paths (basic implementation)
    private static String resolvePath(String path) {
        if (path.startsWith("/")) {
            // Absolute path
            return path.endsWith("/") ? path.substring(1) : path.substring(1) + "/";
        } else if (path.equals("..")) {
            // Parent directory
            int lastSlash = currentDirectory.lastIndexOf("/", currentDirectory.length() - 2);
            if (lastSlash >= 0) {
                return currentDirectory.substring(0, lastSlash + 1);
            }
            return currentDirectory;
        } else if (path.equals(".")) {
            // Current directory
            return currentDirectory;
        } else {
            // Relative path
            return currentDirectory + (path.endsWith("/") ? path : path + "/");
        }
    }

    // Main program
    public static void main(String args[]) {
        if (authUser()) {
            boolean continueShell = true;

            // Create Scanner Object
            Scanner scanner = new Scanner(System.in);

            while (continueShell) {
                // Print shell line
                System.out.print("kelz@javashell:" + (currentDirectory.equals("home/kelz/") ? "~" : currentDirectory) + "$ ");

                // Receive Input
                String userInput = scanner.nextLine();

                // Process userInput
                if (userInput.equals("exit")) // exit -- exits the terminal and closes while loop
                    continueShell = false;
                else
                    processInput(userInput);
            }

            // Close scanner
            scanner.close();
        }
    }
}