import javax.swing.*;
import java.awt.*;
import java.awt.event.*;

public class BankGUI extends JFrame {
    // Reference to our bank resources
    private BankResources bank;

    // GUI Components
    private JPanel loginPanel;
    private JPanel mainPanel;
    private JTextField usernameField;
    private JPasswordField passwordField;
    private JLabel statusLabel;
    private JLabel balanceLabel;

    // Track login attempts
    private int loginAttempts = 0;
    private final int MAX_LOGIN_ATTEMPTS = 3;

    public BankGUI() {
        // Initialize bank resources
        bank = new BankResources();

        // Set up the JFrame
        setTitle("Kelz Online Banking");
        setSize(500, 400);
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setLocationRelativeTo(null); // Center on screen

        // Create panels with CardLayout to switch between login and main banking screens
        CardLayout cardLayout = new CardLayout();
        JPanel contentPanel = new JPanel(cardLayout);

        // Create the login panel
        createLoginPanel();

        // Create the main banking panel (initially hidden)
        createMainPanel();

        // Add panels to the card layout
        contentPanel.add(loginPanel, "login");
        contentPanel.add(mainPanel, "main");

        // Add content panel to frame
        add(contentPanel);

        // Show the login panel first
        cardLayout.show(contentPanel, "login");

        // Make the window visible
        setVisible(true);
    }

    private void createLoginPanel() {
        loginPanel = new JPanel();
        loginPanel.setLayout(new GridBagLayout());
        GridBagConstraints gbc = new GridBagConstraints();
        gbc.insets = new Insets(5, 5, 5, 5);

        // Username components
        JLabel usernameLabel = new JLabel("Username:");
        usernameField = new JTextField(15);

        // Password components
        JLabel passwordLabel = new JLabel("Password:");
        passwordField = new JPasswordField(15);

        // Login button
        JButton loginButton = new JButton("Login");

        // Status label for login errors
        statusLabel = new JLabel("");
        statusLabel.setForeground(Color.RED);

        // Add components to panel with GridBagLayout
        gbc.gridx = 0;
        gbc.gridy = 0;
        gbc.anchor = GridBagConstraints.EAST;
        loginPanel.add(usernameLabel, gbc);

        gbc.gridx = 1;
        gbc.anchor = GridBagConstraints.WEST;
        loginPanel.add(usernameField, gbc);

        gbc.gridx = 0;
        gbc.gridy = 1;
        gbc.anchor = GridBagConstraints.EAST;
        loginPanel.add(passwordLabel, gbc);

        gbc.gridx = 1;
        gbc.anchor = GridBagConstraints.WEST;
        loginPanel.add(passwordField, gbc);

        gbc.gridx = 0;
        gbc.gridy = 2;
        gbc.gridwidth = 2;
        gbc.anchor = GridBagConstraints.CENTER;
        loginPanel.add(loginButton, gbc);

        gbc.gridy = 3;
        loginPanel.add(statusLabel, gbc);

        // Add action listener to login button
        loginButton.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                authenticateUser();
            }
        });

        // Add key listener to handle Enter key in password field
        passwordField.addKeyListener(new KeyAdapter() {
            @Override
            public void keyPressed(KeyEvent e) {
                if (e.getKeyCode() == KeyEvent.VK_ENTER) {
                    authenticateUser();
                }
            }
        });
    }

    private void createMainPanel() {
        mainPanel = new JPanel();
        mainPanel.setLayout(new BorderLayout());

        // Welcome label at the top
        JLabel welcomeLabel = new JLabel("Welcome to Kelz Online Banking", JLabel.CENTER);
        welcomeLabel.setFont(new Font("Arial", Font.BOLD, 18));
        mainPanel.add(welcomeLabel, BorderLayout.NORTH);

        // Balance display
        balanceLabel = new JLabel("Current Balance: $" + bank.getBalance(), JLabel.CENTER);
        balanceLabel.setFont(new Font("Arial", Font.BOLD, 16));

        // Button panel for actions
        JPanel buttonPanel = new JPanel(new GridLayout(5, 1, 10, 10));
        buttonPanel.setBorder(BorderFactory.createEmptyBorder(20, 50, 20, 50));

        JButton depositButton = new JButton("Deposit Funds");
        JButton withdrawButton = new JButton("Withdraw Funds");
        JButton viewBalanceButton = new JButton("Refresh Balance");
        JButton infoButton = new JButton("Personal Information");
        JButton exitButton = new JButton("Exit");

        buttonPanel.add(depositButton);
        buttonPanel.add(withdrawButton);
        buttonPanel.add(viewBalanceButton);
        buttonPanel.add(infoButton);
        buttonPanel.add(exitButton);

        // Add balance and buttons to center
        JPanel centerPanel = new JPanel(new BorderLayout());
        centerPanel.add(balanceLabel, BorderLayout.NORTH);
        centerPanel.add(buttonPanel, BorderLayout.CENTER);

        mainPanel.add(centerPanel, BorderLayout.CENTER);

        // Action listeners for buttons
        depositButton.addActionListener(e -> depositFunds());
        withdrawButton.addActionListener(e -> withdrawFunds());
        viewBalanceButton.addActionListener(e -> updateBalanceDisplay());
        infoButton.addActionListener(e -> displayPersonalInfo());
        exitButton.addActionListener(e -> System.exit(0));
    }

    private void authenticateUser() {
        String username = usernameField.getText();
        String password = new String(passwordField.getPassword());

        if (username.equals("Kelly") && password.equals("pass")) {
            // Switch to main panel on successful login
            Container contentPane = loginPanel.getParent();
            CardLayout cardLayout = (CardLayout) contentPane.getLayout();
            cardLayout.show(contentPane, "main");

            // Reset fields and status
            usernameField.setText("");
            passwordField.setText("");
            statusLabel.setText("");
            loginAttempts = 0;
        } else {
            loginAttempts++;
            if (loginAttempts >= MAX_LOGIN_ATTEMPTS) {
                statusLabel.setText("Max login attempts reached! Try again later.");
                // Disable login button after max attempts
                for (Component comp : loginPanel.getComponents()) {
                    if (comp instanceof JButton) {
                        comp.setEnabled(false);
                    }
                }
            } else {
                statusLabel.setText("Invalid username or password. Attempts: " + loginAttempts + "/" + MAX_LOGIN_ATTEMPTS);
            }
        }
    }

    private void depositFunds() {
        String input = JOptionPane.showInputDialog(this, "Enter amount to deposit:", "Deposit", JOptionPane.PLAIN_MESSAGE);
        try {
            int amount = Integer.parseInt(input);
            if (amount <= 0) {
                JOptionPane.showMessageDialog(this, "Please enter a positive amount.", "Invalid Amount", JOptionPane.ERROR_MESSAGE);
                return;
            }

            int balance = bank.getBalance();
            balance += amount;
            bank.setBalance(balance);

            updateBalanceDisplay();
            JOptionPane.showMessageDialog(this, "Deposit successful! New balance: $" + balance);
        } catch (NumberFormatException e) {
            JOptionPane.showMessageDialog(this, "Please enter a valid number.", "Invalid Input", JOptionPane.ERROR_MESSAGE);
        }
    }

    private void withdrawFunds() {
        String input = JOptionPane.showInputDialog(this, "Enter amount to withdraw:", "Withdrawal", JOptionPane.PLAIN_MESSAGE);
        try {
            int amount = Integer.parseInt(input);
            if (amount <= 0) {
                JOptionPane.showMessageDialog(this, "Please enter a positive amount.", "Invalid Amount", JOptionPane.ERROR_MESSAGE);
                return;
            }

            int balance = bank.getBalance();
            if (balance - amount < 0) {
                JOptionPane.showMessageDialog(this, "Insufficient funds.", "Transaction Failed", JOptionPane.ERROR_MESSAGE);
                return;
            }

            balance -= amount;
            bank.setBalance(balance);

            updateBalanceDisplay();
            JOptionPane.showMessageDialog(this, "Withdrawal successful! New balance: $" + balance);
        } catch (NumberFormatException e) {
            JOptionPane.showMessageDialog(this, "Please enter a valid number.", "Invalid Input", JOptionPane.ERROR_MESSAGE);
        }
    }

    private void updateBalanceDisplay() {
        balanceLabel.setText("Current Balance: $" + bank.getBalance());
    }

    private void displayPersonalInfo() {
        String info = "Account Information:\n\n" +
                "Name: " + bank.getCustomerName() + "\n" +
                "Account Number: " + bank.getAccountNumber() + "\n" +
                "Phone: " + bank.getPhoneNumber();

        JOptionPane.showMessageDialog(this, info, "Personal Information", JOptionPane.INFORMATION_MESSAGE);
    }

    public static void main(String[] args) {
        // Create the GUI on the Event Dispatch Thread
        SwingUtilities.invokeLater(() -> new BankGUI());
    }
}