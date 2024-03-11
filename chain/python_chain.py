import os
import hashlib
import json

# 定义钱包配置文件路径和私钥文件路径
wallet_config_file = "/Users/hcl/work/xiaomi/myGo/mapfeeling/go/chain/wallet_config_python.json"
private_key_file = "/Users/hcl/work/xiaomi/myGo/mapfeeling/go/chain/private_key_python.txt"

# 定义要生成的钱包数量
num_wallets = 10

# 创建钱包配置文件的列表
wallet_config_list = []

# 生成私钥并写入文件
def generate_and_write_private_key():
    private_key = os.urandom(32).hex()  # 生成随机私钥（32字节）并转换为十六进制字符串
    with open(private_key_file, "w") as file:
        file.write(private_key)
    return private_key

# 生成钱包配置文件
def generate_wallet_config(index):
    private_key = generate_and_write_private_key()
    public_key = hashlib.blake2b(hashlib.sha256(bytes.fromhex(private_key)).digest()).hexdigest()  # 计算公钥
    address = hashlib.blake2b(hashlib.sha256(bytes.fromhex(public_key)).digest()).hexdigest()  # 计算地址
    wallet_config = {
        "index": index,
        "privateKey": private_key,
        "publicKey": public_key,
        "address": address
    }
    wallet_config_list.append(wallet_config)
    return wallet_config

# 生成钱包配置文件和私钥文件
for i in range(num_wallets):
    generate_wallet_config(i)
    with open(wallet_config_file, "w") as file:
        json.dump(wallet_config_list, file, indent=4)  # 将钱包配置列表写入JSON文件

print("钱包配置文件已生成：", wallet_config_file)
print("私钥文件已生成：", private_key_file)