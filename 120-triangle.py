class Solution:
    opt = []

    def minimumTotal(self, triangle) -> int:
        self.opt = [[99999 for j in range(len(triangle))] for i in range(len(triangle))]

        minSum = 99999
        for i in range(len(triangle)):
            minSum = min(minSum,self.findOpt(len(triangle)-1,i,triangle))
        return minSum
    
    def findOpt(self,i,j,triangle):
        if i == 0:
            self.opt[0][0] = triangle[0][0]
            return triangle[0][0]
        elif self.opt[i][j] != 99999:
            return self.opt[i][j]
        elif i >= len(triangle) or j >= len(triangle[i]) or i < 0 or j < 0:
            return 99999
        else:
            self.opt[i][j] = triangle[i][j] + min(self.findOpt(i-1,j,triangle),self.findOpt(i-1,j-1,triangle))
            return self.opt[i][j]