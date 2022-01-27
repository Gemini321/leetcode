/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution {
public:
    int kthLargest(TreeNode* root, int k) {
        int cur = 0;
        return search(root, k, cur);
    }
    int search(TreeNode* root, int k, int & cur) {
        int res;
        if(cur < k && root -> right != NULL) res = search(root -> right, k, cur);
        cur += 1;
        if(cur == k) res = root -> val;
        if(cur < k && root -> left != NULL) res = search(root -> left, k, cur);
        return res;
    }
};