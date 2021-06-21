Rails.application.routes.draw do
  # For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html

  resources :announcements
  resources :qiita_posts

  namespace :api_v1, path: '/api/v1', format: 'json' do
    resources :announcements
    resources :qiita_posts
  end
end
