from blockchain import Blockchain

def main() -> None:
    print("BLOCKCHAIN EMULATOR")
    # Initialize a blockchain
    blockchain = Blockchain()

    #DEBUG - print chain
    print(blockchain.chain)

    print("Generating new transaction:")
    blockchain.new_transaction(sender="Kelly", recipient="Haley", amount=5)

    # Debug, show pending transaction
    print(blockchain.current_transactions)

    # Debug, add the transaction to the chain
    # Append new block, with calculated hash from the last block in the chain, and the proof of the last block
    blockchain.chain.append(
        blockchain.new_block(
            len(blockchain.chain),
            blockchain.hash(blockchain.last_block), 
            blockchain.last_block['proof'])
        )

    print(blockchain.chain)


if __name__ == "__main__":
    main()

