;;; GO Emacs setup:
;;; 1) Get https://github.com/dominikh/go-mode.el
;;; 2) Copy Lisp (.el) files to your .emacs dir (or .xemacs).
;;; 3) Change your init.el by adding this at the top (for example
;;;    just after the module header info).
;;; 4) This Lisp is from the second example, tleyden....

;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
;;                      GO mode                                     ;;
;; Check out github go code and copy to ~/.xemacs
;; Ref:  https://github.com/dominikh/go-mode.el
;; http://tleyden.github.io/blog/2014/05/22/configure-emacs-as-a-go-editor-from-scratch/ 
;; 
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
(add-to-list 'load-path "~/.xemacs/")
(require 'go-mode-autoloads)

;; setup GOPATH -- for godoc, may not work
(defun set-exec-path-from-shell-PATH ()
  (let ((path-from-shell (replace-regexp-in-string
                          "[ \t\n]*$"
                          ""
                          (shell-command-to-string "$SHELL --login -i -c 'echo $PATH'"))))
    (setenv "PATH" path-from-shell)
    (setq eshell-path-env path-from-shell) ; for eshell users
    (setq exec-path (split-string path-from-shell path-separator))))

(when window-system (set-exec-path-from-shell-PATH))


;; -- this calls for an absolute path, not sure it is OK
(setenv "GOPATH" ".")
