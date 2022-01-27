#include <vector>
#include <string>
#include <iostream>

class Solution {
public:
    bool is_digit(char c) {
        return (c <= '9' && c >= '0');
    }
    bool is_lower_alpha(char c) {
        return (c <= 'z' && c >= 'a');
    }
    vector<string> split(const string & s, const char delim) {
        vector<string> ans;
        int last_begin = 0, len = s.length(), last_delim = false;
        int delim_cnt;
        for(int i = 0; i < len; i ++) {
            if(s[i] == delim) {
                last_delim = true;
                delim_cnt ++;
            }
            else {
                if(last_delim == true) {
                    ans.push_back(s.substr(last_begin, i - last_begin - delim_cnt));
                    last_begin = i;
                }
                last_delim = false;
                delim_cnt = 0;
            }
        }
        if(s[len - 1] != delim) ans.push_back(s.substr(last_begin, len - last_begin - delim_cnt));
        return ans;
    }
    int countValidWords(string sentence) {
        vector<string> words = split(sentence, ' ');
        int len = words.size();
        int word_cnt = 0;
        for(int i = 0; i < len; i ++) {
            string word = words[i];
            int word_len = word.length(), conn_cnt = 0;
            bool flag = true;
            for(int j = 0; j < word_len; j ++) {
                if(word[j] == '-') {
                    if(j == 0 || j == word_len - 1 || !is_lower_alpha(word[j - 1]) || !is_lower_alpha(word[j + 1])) {
                        flag = false;
                        break;
                    }
                    conn_cnt ++;
                    if(conn_cnt > 1) {
                        flag = false;
                        break;
                    }
                }
                else if(is_lower_alpha(word[j])) continue;
                else if(is_digit(word[j])) {
                    flag = false;
                    break;
                }
                else {
                    if(j != word_len - 1) {
                        flag = false;
                        break;
                    }
                }
            }
            if(flag) word_cnt ++;
            cout << word << " " << flag << endl;
        }
        return word_cnt;
    }
};