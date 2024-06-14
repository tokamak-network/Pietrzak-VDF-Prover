import os
import time
import random
from web3 import Web3


def send_request(web3, contract, private_key, pay, estimated_gas, nonce):

    account = web3.eth.account.from_key(private_key)

    # Generate a tx dictionary
    tx = {
        'nonce': nonce,
        'from': account.address,
        'value': web3.to_wei(pay, 'ether'),
        'gasPrice': web3.eth.gas_price,
        'gas': estimated_gas
    }

    # encoding contract call tx
    function_call = contract.functions.requestRandomWord().build_transaction(tx)

    # sign tx
    signed_tx = web3.eth.account.sign_transaction(function_call, private_key)

    # send tx
    tx_hash = web3.eth.send_raw_transaction(signed_tx.rawTransaction)

    # print tx receipt
    print(f"Transaction hash: {tx_hash.hex()}")

# test function for send; it just prints 'send' on console    
def send():
    print('send')

if __name__=='__main__':

    # Set private key
    private_key = os.getenv('PRIVATE_KEY_1')

    # Set RPC server API
    rpc_url = os.getenv('SEPOLIA_RPC')
    web3 = Web3(Web3.HTTPProvider(rpc_url))

    # Change contract address to the checksum style
    contract_address = web3.to_checksum_address('0xeee4CEF7F7b5938cC73CCb25B34C3eb9DFAdf273')

    contract_abi = [
        {
            "constant": False,
            "inputs": [],
            "name": "requestRandomWord",
            "outputs": [],
            "payable": True,
            "stateMutability": "payable",
            "type": "function"
        }
    ]

    # generate a smart contract instance
    contract = web3.eth.contract(address=contract_address, abi=contract_abi)

    # total number of requests for the test
    number_of_test = 5
    # time duration (sec) for the test
    time_duration = 120
    
    print('** Test setup **')
    print('\t total request: ', number_of_test)
    print('\t time duration: ', time_duration, 'sec')
    
    # Calculate the intervals at which transactions are sent
    intervals = sorted(random.uniform(0, time_duration) for _ in range(number_of_test))
    
    # If there's ETH to pay
    pay = 0.005
    
    # estimate gas
    account = web3.eth.account.from_key(private_key)
    estimated_gas = contract.functions.requestRandomWord().estimate_gas({'from': account.address, 'value': web3.to_wei(pay, 'ether')})
    
    # nonce should be precalculated as this request can be sent before a transaction is confirmed
    # Then, get_transaction_count does not recognize transactions in mempool
    # It causes the transaction replacement
    # Therefore we manually increment nonce one by one
    nonce = web3.eth.get_transaction_count(account.address)
    
    # Record the start time of the test
    start_time = time.time()

    # Initialize the last sent time to zero
    last_sent_time = 0

    # Send all transactions within the for loop
    for interval in intervals:
        # Calculate the delay until the next transaction
        delay = interval - last_sent_time
        time.sleep(delay)
        # send()
        send_request(web3, contract, private_key, pay, estimated_gas, nonce)
        print('nonce: ', nonce)
        nonce += 1
        print(time.time()-start_time)
        # Update last sent time to the current interval
        last_sent_time = interval
        

