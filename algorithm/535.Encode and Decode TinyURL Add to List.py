import base64

class Codec:
    def encode(self, longUrl):
        """Encodes a URL to a shortened URL.
        
        :type longUrl: str
        :rtype: str
        """
        return base64.b64encode(str(longUrl).encode('utf-8'))

    def decode(self, shortUrl):
        """Decodes a shortened URL to its original URL.
        
        :type shortUrl: str
        :rtype: str
        """
        return base64.b64decode(shortUrl)
        

# # Your Codec object will be instantiated and called as such:
# codec = Codec()
# s = codec.encode('1234567')
# ss = codec.decode(s)
# print(s,ss)
