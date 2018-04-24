import ctypes

def GenEcdhKey():
    global EcdhPriKey, EcdhPubKey
    # 载入c模块
    loader = ctypes.cdll.LoadLibrary
    lib = loader("../dll/ecdh_x64.dll")
    # 申请内存
    priKey = bytes(bytearray(2048))                                                         # 存放本地DH私钥
    pubKey = bytes(bytearray(2048))                                                         # 存放本地DH公钥
    lenPri = ctypes.c_int(0)                                                                       # 存放本地DH私钥长度
    lenPub = ctypes.c_int(0)                                                                       # 存放本地DH公钥长度
    # 转成c指针传参
    pri = ctypes.c_char_p(priKey)
    pub = ctypes.c_char_p(pubKey)
    pLenPri = ctypes.pointer(lenPri)
    pLenPub = ctypes.pointer(lenPub)
    # secp224r1 ECC算法
    nid = 713
    # c函数原型:bool GenEcdh(int nid, unsigned char *szPriKey, int *pLenPri, unsigned char *szPubKey, int *pLenPub);
    bRet = lib.GenEcdh(nid, pri, pLenPri, pub, pLenPub)
    if bRet:
        # 从c指针取结果
        EcdhPriKey = priKey[:lenPri.value]
        EcdhPubKey = pubKey[:lenPub.value]
        print(EcdhPriKey)
        print(EcdhPubKey)
        print('我们'.encode('utf-8'))
    return bRet

def main():
    if not GenEcdhKey():
        print('初始化ECC Key失败!')
        return
    print('ok')

if __name__ == '__main__':
    main()
