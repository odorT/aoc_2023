import re
g = []
for x in range(140): g.append(input())
sol = 0
for x in range(len(g)):
    for match in re.finditer('\d+', g[x]):
        try:
            for y in range(*match.span()):
                for i in range(-1, 2):
                    for j in range(-1, 2):
                        assert not (0 <= x+i < len(g) and 0 <= y+j < len(g) and not g[x+i][y+j].isdigit() and g[x+i][y+j] != '.')
        except: 
            print(int(match.group()))
            sol += int(match.group())
print(sol)