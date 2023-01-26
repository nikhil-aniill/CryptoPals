from binascii import hexlify


def repeat_xor(messageq, keyv):
    print("Message is -" + messageq + "and key is " + keyv)
    arr = []
    i = 0
    for ch in messageq:
        arr.append(ord(ch) ^ ord(keyv[i]))
        i = i + 1
        if i == len(keyv):
            i = 0
    return hexlify(bytearray(arr))


if __name__ == "__main__":
    message = "Burning 'em, if you ain't quick and nimble I go crazy when I hear a cymbal"
    key = "ICE"
    cipher = repeatXOR(message, key)
    print(cipher)
